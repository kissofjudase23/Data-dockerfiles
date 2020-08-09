package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // for gorm
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
	Name      string    `gorm:"column:name;type:varchar(255);not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
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

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", user, pass, host, port, dbName)

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	return db
}

func main() {
	ConnectToDB()
	fmt.Println("connect to db success!")
	db := DB.AutoMigrate(&User{})
	if err := db.Error; err != nil {
		log.Fatalf("database migrate error: %v", err)
	}
	fmt.Println("db migration success!")

	// user := User{Name: "Tom"}
	// db.Create(&User{Name: "Tom"})

	user := User{ID: 1}
	db.Model(&user).Update("name", "Tom2")

}
