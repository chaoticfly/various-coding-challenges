package main

import (
	"fmt"
	"regexp"
	"strings"
)

var isLetter = regexp.MustCompile(`^[a-z0-9]+$`).MatchString

func palindrome(str string) bool {
	l := len(str) - 1
	s := 0
	str = strings.ToLower(str)
	for s < l {
		first := byte(0)
		last := byte(0)
		if isLetter(string(str[s])) {
			first = str[s]
		} else {
			s++
		}
		if isLetter(string(str[l])) {
			last = str[l]
		} else {
			l--
		}
		if first != byte(0) && last != byte(0) {

			if first != last {
				return false
			}
			s++
			l--
		}

	}

	return true

}

func main() {
	fmt.Printf("\"a used a car\" is a palindrome? %v\n", palindrome("a used a car"))
	fmt.Printf("\"A man, a plan, a canal: Panama\" is a palindrome? %v\n", palindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("\"Malayalam\" is a palindrome? %v\n", palindrome("Malayalam"))
}
