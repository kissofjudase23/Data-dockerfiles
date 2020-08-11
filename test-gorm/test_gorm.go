package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

// ConnectToDB : connect to database
func ConnectToDB() {
	DB = connectDB()
}

// User : user table
type User struct {
	// mysql.interface
	ID        uint64    `gorm:"column:id;primary_key;auto_increment"`
	Name      string    `gorm:"column:name;type:varchar(255);unique;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime default CURRENT_TIMESTAMP"`
}

// TableName : Specify the User table name
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
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db migration success!")

	users := []User{{ID: 1, Name: "Tom111"}}
	// DB.Create(&users)
	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).Create(&users)

}
