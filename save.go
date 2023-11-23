// save.go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func saveToFile(url, content string) (string, error) {
	dirname := formatDirectoryName(url)
	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		return "", err
	}

	// Save HTML file
	filename := path.Join(dirname, "index.html")
	err = ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return "", err
	}

	// Download and save assets
	assets, err := extractAssets(content, url)
	if err != nil {
		return "", err
	}

	for _, assetURL := range assets {
		err := downloadAndSaveAsset(assetURL, dirname)
		if err != nil {
			fmt.Printf("Error downloading %s: %s\n", assetURL, err)
		}
	}

	return filename, nil
}

func formatDirectoryName(url string) string {
	return strings.ReplaceAll(url, "https://", "")
}

func extractAssets(htmlContent, baseURL string) ([]string, error) {
	var assets []string
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "link" {
			if rel, exists := getAttribute(n, "rel"); exists && rel == "stylesheet" {
				if href, exists := getAttribute(n, "href"); exists {
					assets = append(assets, resolveURL(href, baseURL))
				}
			}
		} else if n.Type == html.ElementNode && n.Data == "img" {
			if src, exists := getAttribute(n, "src"); exists {
				assets = append(assets, resolveURL(src, baseURL))
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}
	extract(doc)

	return assets, nil
}

func getAttribute(n *html.Node, attrName string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == attrName {
			return attr.Val, true
		}
	}
	return "", false
}

func resolveURL(relative, base string) string {
	uri, err := url.Parse(relative)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseURL.ResolveReference(uri)
	return uri.String()
}

func downloadAndSaveAsset(url, dirname string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filename := path.Join(dirname, path.Base(url))
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
