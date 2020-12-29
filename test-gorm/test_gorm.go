package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func getOSEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("can not find os env:%s", key)
	}
	return val
}

func getOSEnvInt(key string) int {
	val := getOSEnv(key)
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}

// DB Handlers
var (
	DB *gorm.DB
)

// PrettyFormat : pretty print for struct
func PrettyFormat(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

// ConnectToDB : connect to database
func ConnectToDB() {
	DB = connectDB()
}

type SMART map[string]interface{}

// Value return json value, implement driver.Valuer interface
// save
func (smart SMART) Value() (driver.Value, error) {
	if len(smart) == 0 {
		return nil, nil
	}
	bytes, err := json.Marshal(smart)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes), err
}

// Scan scan value into Jsonb, implements sql.Scanner interface
// receive
func (j *SMART) Scan(value interface{}) error {
	// null value will not call scan
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := SMART{}
	err := json.Unmarshal(bytes, &result)
	*j = SMART(result)
	return err
}

// GormDataType gorm common data type
func (SMART) GormDataType() string {
	return "smart"
}

// GormDBDataType gorm db data type
func (SMART) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

// User : user table
type User struct {
	// mysql.interface
	Name  string `gorm:"primary_key;type:varchar(255);unique;not null"`
	Smart SMART
}

// TableName : For singular
func (User) TableName() string {
	return "user"
}

func connectDB() *gorm.DB {
	host := getOSEnv("DB_HOST")
	port := getOSEnvInt("DB_PORT")
	user := getOSEnv("DB_USER")
	pass := getOSEnv("DB_PASS")
	dbName := getOSEnv("DB_NAME")
	maxIdle := getOSEnvInt("DB_MAX_IDLE")
	maxOpen := getOSEnvInt("DB_MAX_OPEN")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", user, pass, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqldb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqldb.SetMaxIdleConns(maxIdle)
	sqldb.SetMaxOpenConns(maxOpen)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqldb.SetConnMaxLifetime(d)

	//enable debug mode
	db = db.Debug()

	return db
}

func main() {
	ConnectToDB()
	fmt.Println("connect to db success!")
	tables := []interface{}{
		&User{},
	}
	for _, table := range tables {
		err := DB.AutoMigrate(table)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("db migration success!")

	var err error

	// users := []User{
	// 	{Name: "Tom1", Smart: SMART{"5": 100, "187": 100}},
	// 	{Name: "Tom2", Smart: SMART{}},
	// 	{Name: "Tom3"},
	// }

	// err = DB.CreateInBatches(users, 10).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var usersOut []User
	err = DB.Find(&usersOut).Error
	if err != nil {
		log.Fatal(err)
	}

	PrettyFormat(usersOut)

}
