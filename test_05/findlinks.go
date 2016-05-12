package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
	"net/http"
)

func add(x int, y int) (z int) {
	z = x + y
	return
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	for _, url := range os.Args[1:] {
		findLinks(url)
	}

}

func findLinks(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return false, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	return true, nil
}
