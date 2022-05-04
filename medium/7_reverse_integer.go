package medium

// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Input: x = 123
// Output: 321

func reverse(x int) int {
	const intMax, intMin = 2147483647, -2147483648
	res := 0
	for x != 0 {
		if res > intMax/10 || (res == intMax/10 && x%10 > 7) {
			return 0
		}
		if res < intMin/10 || (res == intMin/10 && x%10 < -8) {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}
	return res
}
