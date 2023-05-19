package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "admin"
	password = "admin"
	hostname = "127.0.0.1:3306"
	dbname   = "db"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func CreateClient() *sql.DB {
	db, err := sql.Open("mysql", dsn())

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	log.Printf("Connected to DB %s successfully\n", dbname)

	return db
}
