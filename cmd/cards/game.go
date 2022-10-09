package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lukas-blaha/quickle/pkg/db"
	"github.com/lukas-blaha/quickle/pkg/game"
)

func PlayCards(conn *sql.DB, lesson string) {
	var t string
	var c, i int
	ask := true

	data, err := db.GetLesson(conn, lesson)
	if err != nil {
		log.Fatal(err)
	}

	game.CleanPrint(" n - next")
	fmt.Println(" p - previous")
	fmt.Println(" e - exit")
	time.Sleep(time.Second * 2)
	// at start print first "term"
	game.CleanPrint(data[i].Term)
	for {
		if ask {
			fmt.Scanln(&t)
		}

		ask = true
		switch {
		// counter "c": even => term, odd => def
		case t == "" && c%2 == 0:
			// if only enter key sent and "term" is active, then print "def"
			game.CleanPrint(data[i].Def)
		case t == "" && c%2 == 1:
			// if only enter key sent and "def" is active, then print "term"
			game.CleanPrint(data[i].Term)
		case t == "n" && i <= len(data)-2:
			// if "n" key is sent, then go to the next entry
			ask = false
			c = 0
			i++
		case t == "p" && i >= 1:
			// if "p" key is sent, then go to the previous entry
			ask = false
			c = 0
			i--
		case t == "e":
			// if "e" key pressed then quit the program
			return
		}

		t = ""
		c++
	}
}
