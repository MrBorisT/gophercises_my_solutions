package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	file, err := os.Open("index.html")
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	var f func(*html.Node)

	links := make([]Link, 0)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			newLink := Link{Text: n.Data}
			for _, a := range n.Attr {
				if a.Key == "href" {
					newLink.Href = a.Val
				}
			}

			links = append(links, newLink)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	for _, l := range links {
		fmt.Printf("Link{\n")
		fmt.Printf("	Href: %v\n", l.Href)
		fmt.Printf("	Text: %v\n", l.Text)
		fmt.Printf("}\n")
	}
}
