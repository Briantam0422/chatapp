package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDBUrl() string {
// 	LoadEnv()
	dbUser := os.Getenv("DATABASE_USER")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbHost, dbPort, dbName)
	return dsn
}

func ConnectDB() *gorm.DB {
	dsn := GetDBUrl()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	CheckError(err)
	fmt.Println("Database has connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	CheckError(err)
	err = sqlDB.Close()
	CheckError(err)
	fmt.Println("Database has closed")
}
