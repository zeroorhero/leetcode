package Str

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := fmt.Sprintf("%d", x)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}
