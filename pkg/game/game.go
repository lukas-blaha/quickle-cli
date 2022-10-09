package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func SelectLesson(choices []string) string {
	sc := bufio.NewScanner(os.Stdin)

	printMessage(choices)
	for sc.Scan() {
		if isNumeric(sc.Text()) {
			n, err := strconv.ParseInt(sc.Text(), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if int(n) <= len(choices) {
				return choices[n-1]
			}
		}

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
	basename := getBasename()
	CleanPrint("Hi, welcome to quickle \"%s\".\n", basename)
	fmt.Printf("\nSelect study set:\n\n")
	for i, l := range lessons {
		fmt.Printf(" %d. %s\n\n", i+1, l)
	}
	fmt.Printf("Select: ")
}
