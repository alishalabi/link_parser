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
	nodes := anchorNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	// dfs(doc, "")
	return nil, nil
}

func anchorNodes(n *html.Node) []*html.Node {
	// If node is an element node, and that element is "a"
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	// Varibale ret = return value
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, anchorNodes(c)...)
	}

	return ret
}

// DFS runs a depth first search of node content
// Returning anchor tags
// func dfs(n *html.Node, padding string) {
// 	// Adding some tag stylings
// 	msg := n.Data 
// 	if n.Type == html.ElementNode {
// 		msg = "<" + msg + ">"
// 	}
// 	fmt.Println(padding, msg)
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding + "   ")
// 	}
// }
