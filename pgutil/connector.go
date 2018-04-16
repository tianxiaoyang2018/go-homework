package pgutil

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once

var db *sql.DB

func GetDB() *sql.DB {

	// 只有一个能获取到锁，其他被阻塞到该方法执行完，然而也不会再调用此方法了
	once.Do(func() {
		db, _ = sql.Open("postgres", "port=5432 user=tianxiaoyang password=txy062151 dbname=txy_db sslmode=disable")
	})

	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
