package main

import (
	"fmt"
	"slices"
)

type page struct {
	url    string
	visits int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	sortedPages := sortedByVisits(pages)
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.visits, page.url)
	}
}

func sortedByVisits(pages map[string]int) []page {
	pageList := make([]page, 0, len(pages))
	for url, visits := range pages {
		pageList = append(pageList, page{url: url, visits: visits})
	}
	slices.SortFunc(pageList, func(a, b page) int {
		return b.visits - a.visits
	})
	return pageList
}
