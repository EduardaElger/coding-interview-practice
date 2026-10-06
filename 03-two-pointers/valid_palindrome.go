package twopointers

import "fmt"

func IsPalindrome(s string) bool {
	fmt.Println(s)
	j := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != s[j] {
			return false
		}
		j++
	}

	return true
}
