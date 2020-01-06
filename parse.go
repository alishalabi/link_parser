package link

import (
	"io"
	"fmt"
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

func dfs(n *html.Node, padding string) {
	fmt.Println(padding, n.Data)
	for c := FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"   ")
	}
}
