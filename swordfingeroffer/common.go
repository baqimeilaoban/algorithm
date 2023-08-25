package swordfingeroffer

import "fmt"

func swapInt(x, y int) (int, int) {
	return y, x
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseArray(arr []int) []int {
	var res []int
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
		fmt.Println(res)
	}
	return res
}

/**
						10
			7							13
	5				9			11				18
3		6		8					12
*/
func generateTreeNode() *TreeNode {
	node10 := &TreeNode{Val: 10}
	node7 := &TreeNode{Val: 7}
	node13 := &TreeNode{Val: 13}
	node5 := &TreeNode{Val: 5}
	node9 := &TreeNode{Val: 9}
	node11 := &TreeNode{Val: 11}
	node18 := &TreeNode{Val: 18}
	node3 := &TreeNode{Val: 3}
	node6 := &TreeNode{Val: 6}
	node8 := &TreeNode{Val: 8}
	node12 := &TreeNode{Val: 12}
	node10.Left, node10.Right = node7, node13
	node7.Left, node7.Right = node5, node9
	node5.Left, node5.Right = node3, node6
	node9.Left = node8
	node13.Left, node13.Right = node11, node18
	node11.Right = node12
	return node10
}

// Push 压栈
func Push(node *TreeNode) {

}

// Pop 出栈
func Pop() *TreeNode {
	return nil
}

// GetTop 获取栈顶元素 不会将栈顶元素弹出来
func GetTop() *TreeNode {
	return nil
}

// IsEmptyStack 栈是否是空
func IsEmptyStack() bool {
	return false
}
