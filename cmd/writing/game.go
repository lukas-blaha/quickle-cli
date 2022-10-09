package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/lukas-blaha/quickle/pkg/db"
	"github.com/lukas-blaha/quickle/pkg/game"
)

func PlayWriting(conn *sql.DB, lesson string) {
	var (
		t, lower, ra string
		i, correct   int
	)

	data, err := db.GetLesson(conn, lesson)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if i == len(data) {
			fmt.Printf("\nYou've got %d correct answers out of %d.\n", correct, len(data))
			time.Sleep(time.Second * 2)
			return
		}
		game.CleanPrint("%s\n", data[i].Term)
		fmt.Printf("Type here: ")
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			t = RemoveBlanks(strings.ToLower(sc.Text()))
			break
		}
		lower = RemoveBlanks(strings.ToLower(data[i].Def))
		ra = RemoveBlanks(RemoveAccent(lower))
		if t == lower || t == ra {
			fmt.Printf("\n\nSuccess, \"%s\" is right answer!\n", data[i].Def)
			time.Sleep(time.Second * 2)
			correct++
		} else {
			fmt.Printf("\n\nIncorrect answer!\n")
			fmt.Printf("The correct answer is: \"%s\"\n", data[i].Def)
			time.Sleep(time.Second * 3)
		}
		i++
	}
}
