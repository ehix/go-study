package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	common "xkcd/common"
	model "xkcd/model"
)

func createUrl(n int) string {
	return fmt.Sprintf("https://xkcd.com/%v/info.0.json", n)
}

func main() {
	var fname string
	var limit int
	flag.StringVar(&fname, "f", common.DefaultFname, "Filename for JSON file output")
	flag.IntVar(&limit, "l", 10, "Limit number of entries pulled from xkcd")

	flag.Parse()

	// Default stopping criteria, two 404s in a row will terminate the loop.
	var notFound int
	var records []model.Record
	for i := 0; i < limit && notFound < 2; i++ {
		url := createUrl(i)
		resp, err := http.Get(url)
		common.CheckError(err)

		if resp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "skipping %v: got %v\n", i, resp.StatusCode)
			notFound++
			continue
		}
		var r model.Record
		r.URL = strings.TrimSuffix(url, "info.0.json")

		err = json.NewDecoder(resp.Body).Decode(&r)
		common.CheckError(err)

		// fmt.Printf("%d: %#v\n", i, r)
		resp.Body.Close()
		notFound = 0
		records = append(records, r)
	}

	fmt.Fprintf(os.Stdout, "read %d comics\n", len(records))

	toJson := model.Records{Records: records}
	asJson, err := json.MarshalIndent(toJson, "", "\t")
	common.CheckError(err)

	dir := common.GetDefaultFilepath(fname)
	fmt.Println(dir)
	err = os.WriteFile(dir, asJson, 0644)
	common.CheckError(err)
}
