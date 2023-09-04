package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	// These are struct tags, a raw string in backticks
	// Mostly used for doing conversion of data in the program to a external format.
	// JSON, XML, PB, SQL, etc.
	// Encodes/decodes with reflection.
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"` // if no words, omit and don't add "null"
	// private fields of a struct that are not exported are not encoded
	// blah string `json:"blah"`
}

func main() {
	r := Response{Page: 1, Words: []string{"up", "in", "out"}}
	j, _ := json.Marshal(r) // returns a bytestring
	fmt.Println(string(j))  // json repres
	fmt.Printf("%#v\n", r)  // internal repres

	var r2 Response
	_ = json.Unmarshal(j, &r2) // takes a []byte and pointer to place to put data.
	fmt.Printf("%#v\n", r2)    // decodes back to the internal repres

}
