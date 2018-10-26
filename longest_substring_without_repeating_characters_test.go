package jason_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	assert.Equal(t, "kew", LongestSubstringWithoutRepeat("pwwkew"))
	assert.Equal(t, "b", LongestSubstringWithoutRepeat("b"))
	assert.Equal(t, "ba", LongestSubstringWithoutRepeat("ba"))
	assert.Equal(t, "vdf", LongestSubstringWithoutRepeat("dvdf"))
	assert.Equal(t, "b", LongestSubstringWithoutRepeat("bbbbb"))
	assert.Equal(t, "abc", LongestSubstringWithoutRepeat("abcabcbb"))
	assert.Equal(t, "ba", LongestSubstringWithoutRepeat("aabaa"))
	assert.Equal(t, "ab", LongestSubstringWithoutRepeat("aabaab"))
	assert.Equal(t, "ab!", LongestSubstringWithoutRepeat("aabaab!bb"))
}

func TestLongestSubstringWithout(t *testing.T) {
	assert.Equal(t, "a", LongestSubstringUsingLetter("a", "a"))
	assert.Equal(t, "b", LongestSubstringUsingLetter("ab", "a"))
	assert.Equal(t, "v", LongestSubstringUsingLetter("dv", "d"))
}
