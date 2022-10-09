package main

import (
	"log"

	"github.com/lukas-blaha/quickle/pkg/db"
	"github.com/lukas-blaha/quickle/pkg/game"
)

func main() {
	// connect to DB
	conn := db.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	choices, err := db.GetStudySets(conn)
	if err != nil {
		log.Fatal(err)
	}

	for {
		lesson := game.SelectLesson(choices)
		PlayWriting(conn, lesson)
		// time.Sleep(time.Second * 5)
	}
}
