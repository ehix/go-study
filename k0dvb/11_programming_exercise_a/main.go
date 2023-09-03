package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// var testRawHtml = `
// <!DOCTYPE html>
// <html>
//   <body>
//     <h1>My First Heading</h1>
//       <p>My first paragraph.</p>
//       <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
//       <img src="xxx.jpg" width="104" height="142">
//   </body>
// </html>
// `

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(-1)
	}
}

// func findScriptTags(n *html.Node) {
// 	if n.Type == html.ElementNode && n.Data == "script" {
// 		for _, attr := range n.Attr {
// 			if attr.Key == "src" {
// 				fmt.Printf("Found script tag with src: %s\n", attr.Val)
// 			}
// 		}
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		findScriptTags(c)
// 	}
// }

func visit(n *html.Node, words, pics *int) {
	if n.Type == html.ElementNode && n.Data == "script" {
		// Ignore JavaScript blocks
		return
	} else if n.Type == html.TextNode {
		// Fields turns a string into a slice of words
		*words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		// Dereference and increment the value
		*pics++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Visit child or next sibling, if nil exit
		// This is depth first search
		visit(c, words, pics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int // Accumulators
	visit(doc, &words, &pics)
	return words, pics
}

func getPage() string {
	// URL of the website you want to fetch
	url := "https://www.amazon.com"

	// Send an HTTP GET request to the URL
	response, err := http.Get(url)
	check(err)
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	check(err)
	// Return the HTML content of the website
	return string(body)
}

func main() {
	htmlString := getPage()
	// HTML parsed as a tree-like structure
	doc, err := html.Parse(bytes.NewReader([]byte(htmlString)))
	check(err)

	words, pics := countWordsAndImages(doc)
	fmt.Printf("%d words, %d images\n", words, pics)
}
