package kata

import (
	"strings"
)

type LongestSubstringWithoutRepeatingCharacters interface {
	lengthOfLongestSubstring(height []int) int
}

type LongestSubstring struct{}

func (c LongestSubstring) lengthOfLongestSubstring(s string) int {
	repeat := LongestSubstringWithoutRepeat(s)
	return len(repeat)
}

func LongestSubstringWithoutRepeat(s string) string {
	longestStr := ""
	currentStr := ""

	for _, v := range s {
		letter := string(v)
		if strings.Contains(currentStr, letter) {
			currentStr = LongestSubstringUsingLetter(currentStr, letter)
		}

		if !strings.Contains(currentStr, letter) {
			currentStr += letter
		}

		if len(currentStr) >= len(longestStr) {
			longestStr = currentStr
		}
	}

	return longestStr
}

func LongestSubstringUsingLetter(currentStr string, without string) string {
	currentSubstringWithout := ""

	if (currentStr == without) {
		return currentStr
	}

	for _, iv := range currentStr {
		letter := string(iv)
		if letter == without {
			currentSubstringWithout = ""
		} else {
			currentSubstringWithout += letter
		}
	}

	return currentSubstringWithout
}
