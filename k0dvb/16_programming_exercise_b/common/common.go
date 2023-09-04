package xkcd

import (
	"fmt"
	"os"
	"path/filepath"
)

const DefaultFname = "xkcd.json"
const DefaultDir = "study/k0dvb/16_programming_exercise_b/common"

func GetDefaultFilepath(alt ...string) string {
	fname := DefaultFname
	if len(alt) > 0 {
		fname = alt[0]
	}
	return filepath.Join(os.Getenv("GOPATH"), DefaultDir, fname)
}

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
