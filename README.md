# Go HTML Link Parser
Get all anchor tags from HTML file. Extract the href and text.

### Model
```
type Link struct {
    Href string
    Text string
}

ParseHTMLLinks(filePath) ([]Link, err)
```

### Getting Started
```
// clone
git clone https://github.com/mmason33/go-html-link-parser.git
cd go-html-link-parser

// run
go run cmd/main.go

// flag
go run cmd/main.go -file=./path-to-html.html
```