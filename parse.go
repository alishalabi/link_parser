package link

import (
	"io"
	// "strings"
	"fmt"

	"golang.org/x/net/html"
)

// Link represents all the data available in an html link
type Link struct {
	Href string
	Text string
}

// Parse will take HTML as input, returns slice of links as output
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

// DFS runs a depth first search of node content
func dfs(n *html.Node, padding string) {
	// Adding some tag stylings
	msg := n.Data 
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding + "   ")
	}
}
