package main

import (
	"strings"
)

func CleanInput(text string) []string {

	wordsLower := strings.ToLower(text)
	words := strings.Fields(wordsLower)

	return words
}
