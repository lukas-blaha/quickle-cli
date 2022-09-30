package main

import (
	"fmt"
	"os"
	"os/exec"
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
