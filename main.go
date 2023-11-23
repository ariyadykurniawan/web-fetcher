// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./fetch <url1> <url2> ...")
		os.Exit(1)
	}

	var urls []string
	for _, arg := range os.Args[1:] {
		urls = append(urls, arg)
	}

	for _, url := range urls {
		htmlContent, err := fetch(url)
		if err != nil {
			fmt.Printf("Error fetching %s: %s\n", url, err)
			continue
		}

		dirname, err := saveToFile(url, htmlContent)
		if err != nil {
			fmt.Printf("Error saving to file: %s\n", err)
			continue
		}

		numLinks, numImages := getMetadata(htmlContent)
		printMetadata(url, numLinks, numImages)

		fmt.Printf("Page content saved to: %s\n", dirname)
	}
}
