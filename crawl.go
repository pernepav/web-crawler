package main

import (
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Printf("invalid baseURL: %s", rawBaseURL)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("invalid currentURL: %s", rawCurrentURL)
		return
	}
	if baseURL.Hostname() != currentURL.Hostname() {
		log.Printf("crawler got out of scope: %s", rawCurrentURL)
		return
	}

	normalizedCurrentURL := normalizeURL(rawCurrentURL)
	if _, ok := pages[normalizedCurrentURL]; ok {
		log.Printf("%s already visited", rawCurrentURL)
		pages[normalizedCurrentURL] += 1
		return
	}

	pages[normalizedCurrentURL] = 1

	log.Printf("fetching html from %s", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("unable to fetch html from %s: %v", rawCurrentURL, err)
		return
	}
	log.Printf("html fetched from %s", rawCurrentURL)

	log.Printf("getting urls from the fetched html")
	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		log.Printf("unable to get urls from the fetched html: %v", err)
		return
	}

	log.Printf("found %d urls", len(urls))
	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
