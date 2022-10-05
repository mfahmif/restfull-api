package app

import (
	"database/sql"
	"restful-api/helper"
	"time"
)

func Newdb() *sql.DB {
	db, err := sql.Open("mysql", "root:fahm0ne1@tcp(localhost:3306)/db_go")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
