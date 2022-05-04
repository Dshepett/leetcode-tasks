package easy

// Given an integer array nums sorted in non-decreasing order, remove
// the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same.
// Since it is impossible to change the length of the array in some languages,
// you must instead have the result be placed in the first part of the array nums.
// More formally, if there are k elements after removing the duplicates,
// then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.

// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	i, j := 0, 1
	for j < len(nums) {

		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		i++
		if j >= len(nums) {
			return i
		}
		nums[i] = nums[j]
	}
	return i
}
