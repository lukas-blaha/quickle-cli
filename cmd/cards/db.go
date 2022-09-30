package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var counts int64

func GetData(db *sql.DB, lesson string) ([]obj, error) {
	var data []obj
	var o obj
	rows, err := db.Query(fmt.Sprintf(`SELECT "term", "def" FROM %s`, lesson))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&o.Term, &o.Def)
		if err != nil {
			return nil, err
		}

		data = append(data, o)
	}

	return data, nil
}

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
	// dsn := os.Getenv("DSN")
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
