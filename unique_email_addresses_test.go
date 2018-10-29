package jason_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumUniqueEmails(t *testing.T) {
	assert.Equal(t, 2, NumUniqueEmails([]string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}))
}
