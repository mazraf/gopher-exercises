package links

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link is the type that is used to represent the result.
// The result that is sent back to caller is the slice of Link.
// It has two fields Href for href and Text for text
type Link struct {
	Href string
	Text string
}

// Parse is the only exported function, it returns slice of Link or error.
// First of all it parses the io.Reader into html document using html lib.
// then it call linkNodes with doc, it is used to return all the <a> nodes.
// with each node from the linkNodes  calls the buildLink which parses the links and returns link which
// is appended to links.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var links []Link
	nodes := linkNodes(doc)
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

// text is used to take *html.Node tag and returns the text that is present inside it.
// It uses recursion get data even from inside tags.
// It looks for textNode, if it found then it returns otherwise it goes inside the child tags recursively.
// At last it takes the string and breaks into slice and takes a slice with " " and joins them and returns.
func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var txt string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		txt += text(c) + " "
	}
	return strings.Join(strings.Fields(txt), " ")
}

// buildLink is called from the Parse. It takes all the <a> tag one by one and
// finds the attributes where key is "href", if href is found it adds it to link, with value.
// then it calls the text() with current node and adds the result to current link.Text field.
func buildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = text(n)
	return link
}

//linkNodes is called from the parse, it returns the slice of *htmlNode where n.Data == "a"
// It uses recurssion to go inside the child nodes.
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
