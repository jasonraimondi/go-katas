package jason_kata

import (
	"fmt"
	"strings"
)

type LongestSubstringWithoutRepeatingCharacters interface {
	lengthOfLongestSubstring(height []int) int
}

type LongestSubstring struct{}

func (c LongestSubstring) lengthOfLongestSubstring(s string) int {
	var substring string
	var longestSubstring string

	for _, v := range s {
		if !strings.Contains(substring, string(v)) {
			substring += string(v)
			if len(substring) > len(longestSubstring) {
				longestSubstring = substring
			}
		} else {
			if len(substring) > len(longestSubstring) {
				longestSubstring = substring
			}
			substring = string(v)
		}
	}

	fmt.Println(longestSubstring)

	return len(longestSubstring)
}
