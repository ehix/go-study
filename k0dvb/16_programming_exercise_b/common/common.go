package xkcd

import (
	"fmt"
	"os"
)

const DefaultFname = "xkcd.json"

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(-1)
	}
}

func CheckFile(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}
