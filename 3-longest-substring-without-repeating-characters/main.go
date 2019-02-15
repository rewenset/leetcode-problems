package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// special cases
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	chars := make(map[byte]int)
	var max, start int

	for i := 0; i < len(s); i++ {
		if p, ok := chars[s[i]]; ok && p >= start {
			start = p + 1
		}

		chars[s[i]] = i

		if max < (i - start + 1) {
			max = i - start + 1
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbb"))     // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring(" "))        // 1
	fmt.Println(lengthOfLongestSubstring("au"))       // 2
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
	fmt.Println(lengthOfLongestSubstring("abba"))     // 2
}
