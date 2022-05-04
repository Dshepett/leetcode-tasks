package easy

// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

func twoSum(nums []int, target int) []int {
	sum := map[int]int{}
	for i, num := range nums {
		if val, ok := sum[num]; ok {
			return []int{val, i}
		}
		sum[target-num] = i
	}
	return nil
}
