package main

import "strings"

func hasFilteredWord(s *string) bool {
	filterWords := []string{
		"fck",
		"fuck",
		"Damn",
		// could be more....
	}

	for _, word := range filterWords {
		if strings.Contains(*s, word) {
			return true
		}
	}
	return false
}
