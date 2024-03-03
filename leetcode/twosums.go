package leetcode

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		pivot := i + 1
		for j, k := range nums[pivot:] {
			if v+k == target {
				return []int{i, j + pivot}
			}
		}

	}
	return []int{}
}
