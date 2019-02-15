package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// special cases
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	var chars map[byte]bool
	var m int

	for i := 0; i < len(s); i++ {
		chars = make(map[byte]bool)

		for j := i; j < len(s); j++ {
			if _, ok := chars[s[j]]; ok {
				break
			}

			chars[s[j]] = true
		}

		if m < len(chars) {
			m = len(chars)
		}
	}

	return m
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbb"))     // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring(" "))        // 1
	fmt.Println(lengthOfLongestSubstring("au"))       // 2
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
}
