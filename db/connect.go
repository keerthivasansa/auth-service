package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var _db *sql.DB

func GetQuerier() *Queries {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	if _db != nil {
		return New(_db)
	}
	driver := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("DB_DSN")
	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	_db = db
	return New(_db)
}
