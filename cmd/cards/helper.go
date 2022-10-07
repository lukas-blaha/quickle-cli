package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func CleanPrint(d string) {
	CallClear()
	fmt.Println(d)
}

func CallClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
