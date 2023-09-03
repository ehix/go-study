package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// recreate `cat`.
// go run fileio.go *.txt
// go run fileio.go *.txt > c.txt
func cat_clone() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		file.Close()
	}
}

// go run fileio.go *.txt
func calc_filesize() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := io.ReadAll(file) // returns a slice of bytes
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println("The file has", len(data), "bytes")
		file.Close()
	}
}

// recreate `wc`
func wc_clone() {
	var tlc, twc, tcc int
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()

			lc++                         // count number of lines
			wc += len(strings.Fields(s)) // counts words
			cc += len(s)                 // counts chars
		}
		tlc += lc
		twc += wc
		tcc += cc
		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()
	}
	fmt.Printf("%7d %7d %7d %s\n", tlc, twc, tcc, "total")
}
