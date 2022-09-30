package main

import (
	"fmt"
	"log"
)

type obj struct {
	Term string
	Def  string
}

func main() {
	var t string
	c, i, last := 0, 0, true

	// connect to DB
	conn := ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	data, err := GetData(conn)
	if err != nil {
		log.Fatal(err)
	}

	CleanPrint(data[i].Term)
	for {
		if last {
			fmt.Scanln(&t)
		}
		last = true
		if t == "" && c%2 == 0 {
			CleanPrint(data[i].Def)
		} else if t == "" && c%2 == 1 {
			CleanPrint(data[i].Term)
		} else if t == "n" && i <= len(data)-2 {
			c, last = 0, false
			i++
		} else if t == "p" && i >= 1 {
			c, last = 0, false
			i--
		} else if t == "e" {
			break
		}
		t = ""
		c++
	}
}
