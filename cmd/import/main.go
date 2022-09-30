package main

import (
	"log"

	"github.com/lukas-blaha/quickle/pkg/db"
)

func main() {
	// connect to DB
	conn := db.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
}
