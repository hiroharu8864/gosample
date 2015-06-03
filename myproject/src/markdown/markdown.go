package main

import (
	"fmt"
	"github.com/russross/blackfriday"
)

func main() {
	markdown := []byte(`
# This is a header
* 項目1
* 項目2
* 項目3
* 項目4
* 項目5
  `)

	html := blackfriday.MarkdownBasic(markdown)
	fmt.Println(string(html))
}
