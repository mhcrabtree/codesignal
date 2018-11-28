package challenges

// CheckPalindrome Given the string, check if it is a palindrome.
func CheckPalindrome(inputString string) bool {
	length := len(inputString)

	for i := 0; i < length/2; i++ {
		if inputString[i] != inputString[length-i-1] {
			return false
		}
	}
	return true
}
