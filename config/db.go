package config

import (
	"erukunrukun/models"
	"fmt"
	"os"

	// "gorm.io/driver/mysql" // if you want use mysql database
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres" // if you want use postgres database
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var DB *gorm.DB

// func InitDB() {
// 	// get ENV
// 	gotenv.Load()

// 	var err error
// 	dbName := os.Getenv("DB_NAME")
// 	dbUsername := os.Getenv("DB_USERNAME")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbPort := os.Getenv("DB_PORT")

// 	// connect to mysql db
// 	// dsn := "root:pass@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
// 	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	// connect to postgresdb
// 	dsn := "user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
// 	fmt.Println(dsn)
// 	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})

// 	if err != nil {
// 		panic("Connecting database failed:" + err.Error())
// 	}

// 	// migrate table
// 	// DB.AutoMigrate(&models.User{})
// 	// DB.AutoMigrate(&models.Item{})
// 	// DB.AutoMigrate(&models.TransactionItem{})
// 	DB.AutoMigrate(&models.WargaMaster{})
// }

func InitDB() *gorm.DB {
	// get ENV
	godotenv.Load()

	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	fmt.Println(dbName)

	// connect to mysql db
	// dsn := "root:pass@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// connect to postgresdb
	dsn := "user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	fmt.Printf("\n\n%#v\n\n", db)
	if err != nil {
		panic("Connecting database failed:" + err.Error())
	}
	// db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.WargaMaster{})
	// db.AutoMigrate(&models.Lookup{})
	// db.AutoMigrate(&models.LookupDetail{})
	// db.AutoMigrate(&models.AllAnak{})
	return db
}
