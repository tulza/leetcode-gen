package utils

import (
	"fmt"
	"strings"
)

func ParseURL(url string) (domain string, routes []string, err error) {
	if url == "" {
		return "", nil, fmt.Errorf("invalid url")
	}

	// trim
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")

	result := strings.Split(url, "/")

	domain = result[0]
	routes = result[1:]
	return domain, routes, nil
}

// https://leetcode.com/problems/two-sum/description/
func HasLeetCodeURL(url string) bool {
    return strings.HasPrefix(url, "https://leetcode.com/problems/")
}