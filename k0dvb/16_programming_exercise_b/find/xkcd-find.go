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
	str = regexp.MustCompile(`[A|a]lt[: |-]`).ReplaceAllString(str, "")
	// Remove remaining punctuation
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
	return strings.ToLower(str)
}

func main() {
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
	// Load the JSON file a decode the contents into Record structs
	records := getRecords(fpath)
	// Find the terms within either the title or transcript
	var count int
outer:
	for _, r := range records {
		title := sanitiseString(r.Title)
		transcript := sanitiseString(r.Transcript)
		for _, term := range terms {
			if !(strings.Contains(title, term) || strings.Contains(transcript, term)) {
				continue outer
			}
		}
		count++
		fmt.Printf("%v %v/%v/%4v '%v'\n", r.URL, r.Day, r.Month, r.Year, r.Title)
	}
	fmt.Printf("found %d matches\n", count)
}
