package jason_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	c := LongestSubstring{}

	assert.Equal(t, 1, c.lengthOfLongestSubstring("b"))
	assert.Equal(t, 1, c.lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, c.lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 3, c.lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 3, c.lengthOfLongestSubstring("dvdf"))
}
