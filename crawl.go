package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	if cfg.maxPagesReached() {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("invalid currentURL %s: %v", rawCurrentURL, err)
		return
	}
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		log.Printf("crawler got out of scope: %s", rawCurrentURL)
		return
	}

	normalizedCurrentURL := normalizeURL(rawCurrentURL)
	isFirstVisit := cfg.addPageVisit(normalizedCurrentURL, rawCurrentURL)
	if !isFirstVisit {
		return
	}

	log.Printf("fetching html from %s", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("unable to fetch html from %s: %v", rawCurrentURL, err)
		return
	}
	log.Printf("html fetched from %s", rawCurrentURL)

	log.Printf("getting urls from the fetched html")
	urls, err := getURLsFromHTML(html, cfg.rawBaseURL)
	if err != nil {
		log.Printf("unable to get urls from the fetched html: %v", err)
		return
	}

	log.Printf("found %d urls", len(urls))
	for _, url := range urls {
		cfg.wg.Add(1)
		go func() {
			defer cfg.wg.Done()
			cfg.concurrencyControl <- struct{}{}
			cfg.crawlPage(url)
			<-cfg.concurrencyControl
		}()
	}
}

func (cfg *config) addPageVisit(normalizedURL, rawCurrentURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normalizedURL]; ok {
		log.Printf("%s already visited", rawCurrentURL)
		cfg.pages[normalizedURL] += 1
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) maxPagesReached() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return len(cfg.pages) >= cfg.maxPages
}
