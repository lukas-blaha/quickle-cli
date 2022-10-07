package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var counts int64

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToDB() *sql.DB {
	dsn := fmt.Sprintln("host=localhost port=5432 user=postgres password=postgres dbname=quizlet sslmode=disable")

	for {
		connection, err := OpenDB(dsn)
		if err != nil {
			counts++
		} else {
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		time.Sleep(2 * time.Second)
		continue
	}
}
