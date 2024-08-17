package main

import (
	"log"
	"maps"
	"os"
	"slices"
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

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("unable to open log.txt: %v", err)
		os.Exit(1)
	}
	defer file.Close()
	log.SetOutput(file)

	baseURL := args[0]
	pages := make(map[string]int)

	crawlPage(baseURL, baseURL, pages)

	log.Print("crawling finished")
	for _, page := range slices.Sorted(maps.Keys(pages)) {
		log.Printf("%s: %d", page, pages[page])
	}
}
