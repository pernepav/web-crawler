package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) string {
	url, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}

	path := strings.TrimRight(url.Path, "/")
	return fmt.Sprintf("%s%s", url.Hostname(), path)
}
