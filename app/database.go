package app

import (
	"belajar_belajar_golang_restfull_api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/belajar_golang-restfull_api")

	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)

	db.SetConnMaxIdleTime(10 * time.Second)

	return db
}
