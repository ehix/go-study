package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// count unique words in a file
func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	// give one word at a time
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		word := scan.Text()
		words[strings.ToLower(word)]++
	}

	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var skv []kv
	for k, v := range words {
		skv = append(skv, kv{k, v})
	}

	// func literal, can create anywhere you can put anywhere else
	// a lamda, anon, etc
	sort.Slice(skv, func(i, j int) bool {
		// sort function, how to order
		// get most frequent words first
		return skv[i].val > skv[j].val
	})

	for _, s := range skv[:5] {
		// show top 5
		fmt.Printf("%6v x %v\n", s.key, s.val)
	}
}
