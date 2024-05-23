package app

//postgres
import (
	"database/sql"
	"fmt"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Connect to the database
func NewDB() *sql.DB {

	env := godotenv.Load()
	helper.PanicIfError(env)

	user := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_DATABASE")
	sslmode := os.Getenv("DB_SSLMODE")

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s", user, dbname, sslmode))
	helper.PanicIfError(err)
	
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	//migrate -database "postgres://postgres:123456@localhost:5432/belajar_golang_restful_api?sslmode=disable" -path db/migrations up / down / force {version before dirty} / version
	//migrate create -ext sql -dir db/migrations -tz ASIA/JAKARTA create_table_tests
}