package pgutil

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once

var db *sql.DB

func createDB() {
	var err error
	db, err = sql.Open("postgres", "port=5432 user=tianxiaoyang password=txy062151 dbname=txy_db sslmode=disable")
	CheckErr(err)
}

func GetDB() *sql.DB {
	if db == nil {
		once.Do(createDB)
	}
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
