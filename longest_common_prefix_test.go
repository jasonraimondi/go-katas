package jason_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "pw", LengthOfLongestCommonPrefix([]string{"pwwkew", "pwk"}))
	assert.Equal(t, "", LengthOfLongestCommonPrefix([]string{"dog","racecar","car"}))
	assert.Equal(t, "fl", LengthOfLongestCommonPrefix([]string{"flower","flow","flight"}))
	assert.Equal(t, "reg", LengthOfLongestCommonPrefix([]string{"register","region","reglobal"}))
	assert.Equal(t, "", LengthOfLongestCommonPrefix([]string{"Re","region","reglobal"}))
	assert.Equal(t, "", LengthOfLongestCommonPrefix([]string{"aca","cba"}))
}
