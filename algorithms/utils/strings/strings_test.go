package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	assert.True(t, Palindrome(""))
	assert.True(t, Palindrome("a"))
	assert.True(t, Palindrome("aa"))
	assert.True(t, Palindrome("aaa"))
	assert.False(t, Palindrome("ab"))
	assert.True(t, Palindrome("aba"))
	assert.False(t, Palindrome("abab"))
}
