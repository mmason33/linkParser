package main

import (
	"flag"
	"fmt"

	"github.com/mmason33/linkParser"
)

// https://godoc.org/golang.org/x/net/html#example-Parse
// https://github.com/gophercises/link

func main() {
	fileFlag := flag.String("file", "./ex1.html", "Path to the html file")
	flag.Parse()
	links, err := linkParser.ParseHTMLLinks(*fileFlag)
	fmt.Println(links, err)
}
