package swordfingeroffer

import (
	"fmt"
	"testing"
)

func TestWeek1(t *testing.T) {
	nums := "We are happy."
	res := replaceSpaces(nums)
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
		// i != nums[i] 主要是为了排除无法交换的场景，如 0 == num[0]，此场景该进入下一个步骤
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

/**
在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维 数组和一个整数，判断数组中是否含有该整数
数据范围：
二维数组中元素个数范围 [0,1000]
示例：
输入数组：
[
  [1,2,8,9]，
  [2,4,9,12]，
  [4,7,10,13]，
  [6,8,11,15]
]
如果输入查找数值为7，则返回true，
如果输入查找数值为5，则返回false
*/
func searchArray(array [][]int, target int) bool {
	// 直接对比左上角的数，例如9，比9大就会把行减一，比9小就会把列减一
	length := len(array)
	// 边界条件
	if length == 0 || len(array[0]) == 0 {
		return false
	}
	i, j := 0, len(array[0])-1
	for i < length && j >= 0 {
		if array[i][j] == target {
			return true
		} else if array[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

/**
请实现一个函数，把字符串中的每个空格替换成"%20"。
数据范围
0≤ 输入字符串的长度 ≤1000。
注意输出字符串的长度可能大于 1000。
样例：
输入："We are happy."
输出："We%20are%20happy."
*/
func replaceSpaces(str string) string {
	var res string
	for _, c := range str {
		if c == ' ' {
			res = fmt.Sprintf("%s%s", res, "%20")
		} else {
			res = fmt.Sprintf("%s%c", res, c)
		}
	}
	return res
}

/**
输入一个链表的头结点，按照 从尾到头 的顺序返回节点的值。
返回的结果用数组存储。
数据范围
0≤ 链表长度 ≤1000
样例：
输入：[2, 3, 5]
返回：[5, 3, 2]
*/
func printListReversion(head *ListNode) []int {
	var res []int
	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return reverseArray(res)
}
