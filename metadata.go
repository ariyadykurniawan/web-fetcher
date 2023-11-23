// metadata.go
package main

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func getMetadata(htmlContent string) (int, int) {
	doc, _ := html.Parse(strings.NewReader(htmlContent))
	return countLinks(doc), countImages(doc)
}

func countLinks(n *html.Node) int {
	if n.Type == html.ElementNode && n.Data == "a" {
		return 1
	}

	count := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count += countLinks(c)
	}
	return count
}

func countImages(n *html.Node) int {
	if n.Type == html.ElementNode && n.Data == "img" {
		return 1
	}

	count := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count += countImages(c)
	}
	return count
}

func printMetadata(url string, numLinks, numImages int) {
	fmt.Printf("site: %s\n", url)
	fmt.Printf("num_links: %d\n", numLinks)
	fmt.Printf("images: %d\n", numImages)
	fmt.Printf("last_fetch: %s\n", time.Now().UTC().Format("Mon Jan 02 2006 15:04:05 UTC"))
}
