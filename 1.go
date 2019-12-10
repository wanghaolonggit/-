/*
该题描述：

给定一个长度为n的整数数组nums，数组中所有的数字都在 0∼n−1的范围内。

数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。

请找出数组中任意一个重复的数字。

注意：如果某些数字不在 0∼n−1 的范围内，或数组中不包含重复数字，则返回 -1；

*/

func duplicateInArray(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return -1
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		//移动到i 的位置
		for nums[i] != i {
			if nums[nums[i]] != nums[i] {
				nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
			} else {
				return nums[i]
			}
		}
	}
	return -1
}
