package main

import (
	"fmt"
	"strings"
)

func AligningMinus(text string) string {
	if !strings.ContainsRune(text, '-') {
		return " " + text
	}
	return text
}

func clearFromTop() {
	//from bottom
	fmt.Printf("\x1bc")
}
