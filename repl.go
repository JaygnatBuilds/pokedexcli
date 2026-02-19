package main

import (
	"strings"
)

func cleanInput(text string) []string {

	words := strings.Fields(text)

	return words
}
