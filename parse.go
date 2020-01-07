package main

import (
	"io"
	"strings"
	// "fmt"

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
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLinkObject(node))
	}
	// dfs(doc, "")
	return links, nil
}

func buildLinkObject(n *html.Node) Link {
	// Varibale ret = return value
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = getText(n)
	return ret
}

// getText extracts all text from a given node
func getText(n *html.Node) string {
	// Base Case: node is a text node -> return
	if n.Type == html.TextNode {
		return n.Data
	}
	// Base Case: node is not an element node -> skip
	if n.Type != html.ElementNode {
		return ""
	}
	// Varibale ret = return value
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c) + " "
	}

	// Formating (using string package), converting whitespace to slice,
	// joining slice into string
	return strings.Join(strings.Fields(ret), " ")
}

// anchorNodes recursively searches all nodes and
// returns all anchor nodes
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
