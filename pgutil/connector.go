package pgutil

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PgConnect() *sql.DB {
	fmt.Println("------ 连接数据库 ------")
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "port=5432 user=tianxiaoyang password=txy062151 dbname=txy_db sslmode=disable")
	CheckErr(err)
	return db
}
func PgClose(db *sql.DB) {
	db.Close()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
