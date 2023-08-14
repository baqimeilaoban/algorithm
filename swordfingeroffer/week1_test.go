package swordfingeroffer

import (
	"fmt"
	"testing"
)

func TestWeek1(t *testing.T) {
	nums := []int{2, 3, 5, 4, 3, 2, 6, 7}
	res := duplicateInArray02(nums)
	fmt.Println(res)
}

/*
题1：找出数组中重复的数字
给定一个长度为 n 的整数数组 nums，数组中所有的数字都在 0∼n−1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。
注意：如果某些数字不在 0∼n−1的范围内，或数组中不包含重复数字，则返回 -1；
数据范围
0≤n≤1000

样例：
给定 nums = [2, 3, 5, 4, 3, 2, 6, 7]。
返回 2 或 3。

时间复杂度 o(n) 空间复杂度 0(1)
*/
func duplicateInArray(nums []int) int {
	size := len(nums)
	// 判断边界 0~n-1
	for _, num := range nums {
		if num < 0 || num >= size {
			return -1
		}
	}
	for i := 0; i < size; i++ {
		// i != nums[i] 主要是为了排除无法交换的场景，如 0 == num[0]，此场景该进入下一步骤
		for i != nums[i] && nums[nums[i]] != nums[i] {
			nums[i], nums[nums[i]] = swapInt(nums[i], nums[nums[i]])
		}
		if nums[i] != i && nums[nums[i]] == nums[i] {
			return nums[i]
		}
	}
	return -1
}

/*
题2：不修改数组找出重复的数字
给定一个长度为 n+1 的数组nums，数组中所有的数均在 1∼n 的范围内，其中 n≥1
请找出数组中任意一个重复的数，但不能修改输入的数组。
数据范围
1≤n≤1000

示例：
给定 nums = [2, 3, 5, 4, 3, 2, 6, 7]。
返回 2 或 3。
*/
func duplicateInArray02(nums []int) int {
	// 二分法+抽屉原理
	l, r := 1, len(nums)-1
	for l < r {
		mid := (l + r) >> 1 // [l, mid] [mid+1, r]
		// 定义在左边数的个数
		s := 0
		for _, num := range nums {
			if num >= l && num <= mid {
				s++
			}
		}
		if s > mid-l+1 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func searchArray() {

}
