package strings

func Palindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
