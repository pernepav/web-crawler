package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	hrefs := processHTMLNode(doc)

	for i := 0; i < len(hrefs); i++ {
		if isRelative(hrefs[i]) {
			if strings.HasPrefix(hrefs[i], "/") {
				hrefs[i] = rawBaseURL + hrefs[i]
			} else {
				hrefs[i] = rawBaseURL + "/" + hrefs[i]
			}
			continue
		}
	}

	return hrefs, nil
}

func processHTMLNode(node *html.Node) []string {
	var hrefs []string
	if node.Type == html.ElementNode && node.Data == "a" && len(node.Attr) > 0 {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				hrefs = append(hrefs, attr.Val)
				break
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		hrefs = append(hrefs, processHTMLNode(child)...)
	}

	return hrefs
}

func isRelative(href string) bool {
	url, _ := url.Parse(href)
	return !url.IsAbs()
}
