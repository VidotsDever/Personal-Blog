package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:weichuang@/blog?charset=utf8")
	if err != nil {
		panic(err)
	}
}
