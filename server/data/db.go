package data

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	connectionDB := GetConnection()
	db = connectionDB
}

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}

	//get connection attr from .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASS")
	// dbname := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")
	// sslMode := os.Getenv("SSL_MODE")
	// timezone := os.Getenv("DB_TIMEZONE")

	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, pass, dbname, port, sslMode, timezone)
	log.Print("Connecting to database...")
	dsn := "root:@tcp(127.0.0.1:3306)/task_management_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Connecting to database...")
		panic("cannot connect to database")
	}
	log.Print("Database connected!")

	return db
}

func Migrate() {
	log.Print("Migrating data...")
	db.AutoMigrate(&User{}, &Task{})
	log.Print("Migrating success")
}
