package main

import (
	"log"
)

type obj struct {
	Term string
	Def  string
}

func main() {
	// connect to DB
	conn := ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	for {
		lesson := SelectLesson()
		PlayCards(conn, lesson)
	}
}
