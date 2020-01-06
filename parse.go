package link_parser

import (
	"io"
)

// Link represents all the data available in an html link
type Link struct {
	Href string
	Text string
}

// Parse will take HTML as input, returns slice of links as output
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
