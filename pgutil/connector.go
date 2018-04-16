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
		// 判断db为空的线程进来，可能进来了多个线程，只有一个能获取到锁，其他被阻塞到该方法执行完，然而也不会再调用此方法了
		once.Do(createDB)
	}
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
