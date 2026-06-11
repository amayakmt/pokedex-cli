package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	result := strings.Fields(lower)
	return result
}
