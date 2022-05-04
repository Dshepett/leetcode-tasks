package easy

// Given an integer x, return true if x is palindrome integer.

// Input: x = 121
// Output: true

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	z := x
	val := 0
	for z > 0 {
		val = val*10 + z%10
		z /= 10
	}
	return val == x
}
