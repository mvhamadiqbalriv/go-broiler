package app

import (
	"fmt"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to the database using GORM
func NewDB() *gorm.DB {
	host := helper.GetEnv("DB_HOST")
	port := helper.StringToInt(helper.GetEnv("DB_PORT"))
	user := helper.GetEnv("DB_USERNAME")
	password := helper.GetEnv("DB_PASSWORD")
	dbname := helper.GetEnv("DB_DATABASE")
	sslmode := helper.GetEnv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta", host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	sqlDB, err := db.DB()
	helper.PanicIfError(err)

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
