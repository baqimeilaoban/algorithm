package swordfingeroffer

import "fmt"

func swapInt(x, y int) (int, int) {
	return y, x
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseArray(arr []int) []int {
	var res []int
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
		fmt.Println(res)
	}
	return res
}
