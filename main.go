package main

import (
	"fmt"

	hashmap "coding-studing/02-hash-map"
	twopointers "coding-studing/03-two-pointers"
	slidingwindow "coding-studing/04-sliding-window"
)

func main() {
	fmt.Println("Two Sum:")
	fmt.Println(hashmap.TwoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println("Valid Anagram:")
	fmt.Println(hashmap.IsAnagram("anagram", "nagaram"))

	fmt.Println("Valid Palindrome:")
	fmt.Println(twopointers.IsPalindrome("racecar"))

	fmt.Println("Maximum Subarray Sum:")
	fmt.Println(slidingwindow.MaxSubarraySum([]int{2, 1, 5, 1, 3, 2}, 3))
}
