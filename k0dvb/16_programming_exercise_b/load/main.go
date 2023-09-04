package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	common "ex16b/common"
	model "ex16b/model"
)

func createUrl(n int) string {
	return fmt.Sprintf("https://xkcd.com/%v/info.0.json", n)
}

func main() {
	var fname string
	flag.StringVar(&fname, "f", common.DefaultFname, "Filename for JSON file output")
	flag.Parse()

	limit := 5
	var records []model.Record
	for i := 0; i < limit; i++ {
		url := createUrl(i)
		resp, err := http.Get(url)
		common.CheckError(err)

		if resp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "skipping %v: got %v\n", i, resp.StatusCode)
			continue
		}
		var r model.Record
		r.URL = url

		err = json.NewDecoder(resp.Body).Decode(&r)
		common.CheckError(err)

		// fmt.Printf("%d: %#v\n", i, r)
		resp.Body.Close()
		records = append(records, r)
	}

	fmt.Fprintf(os.Stdout, "read %d comics\n", len(records))

	asJson, err := json.MarshalIndent(records, "", "\t")
	common.CheckError(err)

	dir := fmt.Sprintf("./%s", fname)
	fmt.Println(dir)
	err = os.WriteFile(dir, asJson, 0644)
	common.CheckError(err)
}
