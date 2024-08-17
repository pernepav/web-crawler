package main

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 3 {
		log.Println("too many arguments provided")
		os.Exit(1)
	}

	log.SetOutput(os.Stdout)

	baseURL, err := url.Parse(args[0])
	if err != nil {
		log.Fatalf("invalid base url %s: %v", args[0], err)
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("invalid max concurrency %s: %v", args[1], err)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatalf("invalid max pages %s: %v", args[2], err)
		os.Exit(1)
	}

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		rawBaseURL:         args[0],
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.crawlPage(args[0])

	cfg.wg.Wait()

	log.Print("crawling finished")

	printReport(cfg.pages, cfg.rawBaseURL)
}
