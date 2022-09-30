package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func SelectLesson() string {
	choices := []string{"lesson01"}

	sc := bufio.NewScanner(os.Stdin)

	printMessage(choices)
	for sc.Scan() {
		for _, l := range choices {
			if sc.Text() == l {
				return l
			}
		}
		CleanPrint("Incorrect option, please try again...")
		time.Sleep(time.Second * 1)
		printMessage(choices)
	}
	return ""
}

func printMessage(lessons []string) {
	CleanPrint("Hi, welcome to quickle cards.")
	fmt.Printf("\nSelect study set:\n")
	for _, l := range lessons {
		fmt.Printf(" - %s\n\nSelect: ", l)
	}
}

func PlayCards(conn *sql.DB, lesson string) {
	var t string
	var c, i int
	ask := true

	data, err := GetData(conn, lesson)
	if err != nil {
		log.Fatal(err)
	}

	CleanPrint(" n - next")
	fmt.Println(" p - previous")
	fmt.Println(" e - exit")
	time.Sleep(time.Second * 2)
	// at start print first "term"
	CleanPrint(data[i].Term)
	for {
		if ask {
			fmt.Scanln(&t)
		}

		ask = true
		switch {
		// counter "c": even => term, odd => def
		case t == "" && c%2 == 0:
			// if only enter key sent and "term" is active, then print "def"
			CleanPrint(data[i].Def)
		case t == "" && c%2 == 1:
			// if only enter key sent and "def" is active, then print "term"
			CleanPrint(data[i].Term)
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
