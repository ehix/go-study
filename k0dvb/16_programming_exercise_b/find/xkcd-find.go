package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	common "xkcd/common"
	model "xkcd/model"
)

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	unique := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			unique = append(unique, sanitiseString(item))
		}
	}
	return unique
}

func getRecords(fpath string) []model.Record {
	jsonFile, err := os.Open(fpath)
	common.CheckError(err)

	defer jsonFile.Close()

	var records model.Records
	err = json.NewDecoder(jsonFile).Decode(&records)
	common.CheckError(err)
	return records.Records
}

func sanitiseString(str string) string {
	// Will preserve ints but not floats, e.g. "2.5" becomes "25".
	// Make new lines spaces
	str = regexp.MustCompile(`\r?\n`).ReplaceAllString(str, " ")
	// Remove "alt/Alt: "
	str = regexp.MustCompile(`[A|a]lt[:|-]`).ReplaceAllString(str, "")
	// Remove remaining punctuation
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
	return strings.ToLower(str)
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

	// Remove any duplicate search terms passed
	terms := removeDuplicateStr(os.Args[2:])

	records := getRecords(fpath)

	var return_records model.Records
	for _, r := range records {

		words := make(map[string]int)
		for _, word := range strings.Fields(sanitiseString(r.Transcript)) {
			words[word]++
		}

		for _, term := range terms {
			_, ok := words[term]
			if ok {
				return_records.AddRecord(r)
				break
			}
		}
	}
	fmt.Printf("found %d matches\n", len(return_records.Records))
	for i, r := range return_records.Records {
		fmt.Printf("%3d: %v, %2v/%v/%4v, %v\n", i+1, r.URL, r.Day, r.Month, r.Year, r.Title)
	}
}
