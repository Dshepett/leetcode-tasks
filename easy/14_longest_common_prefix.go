package easy

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string ""

// Input: strs = ["flower","flow","flight"]
// Output: "fl"

func longestCommonPrefix(strs []string) string {
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return res
			}
			if strs[j][i] != strs[0][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res
}
