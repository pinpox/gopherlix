package main

import (
	"strings"
)

// Make \r\n readable in test output
func replaceCRLF(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "\r", "\\r"), "\n", "\\n")
}
