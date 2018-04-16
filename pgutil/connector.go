package pgutil

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("postgres", "port=5432 user=tianxiaoyang password=txy062151 dbname=txy_db sslmode=disable")
		CheckErr(err)
	}
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
