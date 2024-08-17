package main

import (
	"log"
	"maps"
	"net/url"
	"os"
	"slices"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		log.Println("too many arguments provided")
		os.Exit(1)
	}

	// file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatalf("unable to open log.txt: %v", err)
	// 	os.Exit(1)
	// }
	// defer file.Close()
	// log.SetOutput(file)

	baseURL, err := url.Parse(args[0])
	if err != nil {
		log.Fatalf("invalid base url %s: %v", args[0], err)
		os.Exit(1)
	}

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		rawBaseURL:         args[0],
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
	}

	cfg.crawlPage(args[0])

	cfg.wg.Wait()

	log.Print("crawling finished")

	for _, page := range slices.Sorted(maps.Keys(cfg.pages)) {
		log.Printf("%s: %d", page, cfg.pages[page])
	}
}
