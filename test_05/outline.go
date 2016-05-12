package main

import (
	"fmt"
	"strings"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func add1(r rune) rune {
	return r + 1
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			fmt.Println("multiple title elements")
		default:
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func main() {
	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, "abc123"))
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get: %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "getting %s: %s", url, resp.Status)
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		fmt.Fprintf(os.Stderr, "%s has type %s,not text/html", url, ct)
	}
	body := strings.NewReader("<html><body><title>The Go Programming 1</title><title>The Go Programming 2</title></body></html>")
	//doc, err := html.Parse(resp.Body)

	doc, err := html.Parse(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	//forEachNode(doc, startElement, endElement)
	title, err := soleTitle(doc)
	fmt.Println("title:", title)
}
