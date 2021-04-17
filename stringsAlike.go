package main

import "strings"

func main() {

}

func halvesAreAlike(s string) bool {

}

func isVowel(s string) bool {
	var lower = strings.ToLower(s)
	switch lower {
	case "a":
	case "e":
	case "i":
	case "o":
	case "u":
		return true
	}

	return false

}
