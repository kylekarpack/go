package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("test"))
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(str string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	str = reg.ReplaceAllString(str, "")
	str = strings.ToLower(str)
	var len = len(str)
	for i := 0; i < len-1; i++ {
		var e = len - i - 1
		if str[i] != str[e] {
			return false
		}
	}
	return true
}
