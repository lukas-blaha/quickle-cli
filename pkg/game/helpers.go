package game

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func CleanPrint(d string, args ...any) {
	CallClear()
	fmt.Printf(d, args...)
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

func getBasename() string {
	n := strings.Split(os.Args[0], "/")
	basename := n[len(n)-1]
	return basename
}
