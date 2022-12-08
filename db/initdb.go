package db

import (
	"database/sql"

	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local",

		"root", "test", "34.81.56.0:3306", "new_schema", "utf8")

	if conn, err := sql.Open("mysql", dsn); err != nil {

		panic(err.Error())

	} else {

		conn.SetConnMaxLifetime(7 * time.Second)

		conn.SetMaxOpenConns(10)

		conn.SetMaxIdleConns(10)

		return conn

	}

}
