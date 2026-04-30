package algoritmosexercicios

// entrada é ASCII only
func IsPalindrome(s string) bool {
	palindromeBuffer := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		v := s[i]
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			palindromeBuffer = append(palindromeBuffer, v)
		} else if v >= 'A' && v <= 'Z' {
			palindromeBuffer = append(palindromeBuffer, v+'a'-'A')
		}
	}
	left := 0
	right := len(palindromeBuffer) - 1

	for left < right {
		if palindromeBuffer[left] != palindromeBuffer[right] {
			return false
		}
		left++
		right--
	}

	return true

}
