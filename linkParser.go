package linkParser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

//Link stores a link with text in it
type Link struct {
	Href string
	Text string
}

//Parse parses the HTML file and returns Links
func ParseHTMLLinks(filePath string) ([]Link, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	r := strings.NewReader(string(file))
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return getLinks(doc), nil

}

func getLinks(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				txt := getText(n)
				links = append(links, Link{a.Val, txt})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		exLinks := getLinks(c)
		links = append(links, exLinks...)
	}
	return links
}

func getText(n *html.Node) string {
	var text string
	if n.Type != html.ElementNode && n.Data != "a" && n.Type != html.CommentNode {
		text = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}
	return strings.Trim(text, "\n")
}
