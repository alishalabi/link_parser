package main

import (
	"fmt"
	"strings"


	"github.com/alishalabi/link_parser"
)

var exampleHtml = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
