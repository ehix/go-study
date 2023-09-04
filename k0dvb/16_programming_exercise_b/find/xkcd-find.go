package main

import (
	"fmt"
	"os"

	common "xkcd/common"
)

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	unique := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}

func main() {
	// <!> sort out arg validation:
	// largs := len(os.Args)
	// switch largs {
	// case largs < 2:
	// 	fmt.Fprintln("Requires filename and search terms be passed")
	// case largs < 3:
	// }

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stdout, "No search terms passed")
		os.Exit(0)
	}

	fpath := common.GetDefaultFilepath(os.Args[1])
	if !common.CheckFile(fpath) {
		fmt.Fprintf(os.Stderr, "Cannot find file given '%s'\n", os.Args[1])
		os.Exit(0)
	}

	terms := removeDuplicateStr(os.Args[2:])
	fmt.Printf("%v\n", terms)

	// load the file
	// Open our jsonFile
	jsonFile, err := os.Open(fpath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}
