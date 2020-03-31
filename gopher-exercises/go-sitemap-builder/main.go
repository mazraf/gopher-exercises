package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/farzamalam/gopher-exercises/go-link-parser/links"
)

// xmlns is the namespace associated with the sitemap xml.
const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

// loc contains the url.
type loc struct {
	Value string `xml:"loc"`
}

// urlset is the root element of our xml
type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

// main first of takes two flags urlFlag for userentered url and maxDepth.
// It calls the bfs() with urlFlag and maxDepth and gets the result in hrefs.
// It then takes the hrefs and encodes them in the xml form using encoding/xml
func main() {
	// Created a flag url with default "https://www.gophercises.com"
	urlFlag := flag.String("url", "https://gophercises.com", "the url to make sitemap for.")
	maxDepth := flag.Int("depth", 3, "maximum depth of search")
	flag.Parse()
	hrefs := bfs(*urlFlag, *maxDepth)

	toXml := urlset{
		Xmlns: xmlns,
	}

	for _, href := range hrefs {
		toXml.Urls = append(toXml.Urls, loc{href})
	}
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", " ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
}

// bfs is used to iteratively go inside each link and call get method on it.
// It uses Breadth-first algorithm. It calls the get()
func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for u := range seen {
		ret = append(ret, u)
	}
	return ret
}

// get takes the url and returns a filtered slice of links on the same page.
// It calls the url with GET http method and returns by sending html document to href with base url.
func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	return filter(base, href(resp.Body, base))
}

// href is used to call links.Parse with html doc and return a result that contans string slice
// of all the urls with base associated with it.
// It only takes two form of url i.e "/" or "http", it ignores all those like #site etc
func href(r io.Reader, base string) []string {
	var res []string
	links, _ := links.Parse(r)
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			res = append(res, base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			res = append(res, link.Href)
		}
	}
	return res
}

// filter takes two arguments, base and slice of hrefs
// it returns only those hrefs that contains base.
// this is used to ensure that we are building map for internal links only.
func filter(base string, hrefs []string) []string {
	var filterHref []string
	for _, href := range hrefs {
		if strings.HasPrefix(href, base) {
			filterHref = append(filterHref, href)
		}
	}
	return filterHref
}
