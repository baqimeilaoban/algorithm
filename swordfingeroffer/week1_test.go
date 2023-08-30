package swordfingeroffer

import (
	"fmt"
	"testing"
)

func TestWeek1(t *testing.T) {
	node := generateTreeNode()
	postOrder(node)
	fmt.Println(postOrderRes)
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

/**
输入一棵二叉树前序遍历和中序遍历的结果，请重建该二叉树。
注意:
二叉树中每个节点的值都互不相同；
输入的前序遍历和中序遍历一定合法；
数据范围
树中节点数量范围 [0,100]。
样例：
给定：
前序遍历是：[3, 9, 20, 15, 7] 根左右
中序遍历是：[9, 3, 15, 20, 7] 左根右
返回：[3, 9, 20, null, null, 15, 7, null, null, null, null]
返回的二叉树如下所示：
    3
   / \
  9  20
    /  \
   15   7
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 前序遍历的形式：[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
	// 中序遍历的形式：[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
	if len(preorder) == 0 {
		return nil
	}
	// 生成根节点
	root := &TreeNode{preorder[0], nil, nil}
	var x int
	for i := 0; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			x = i
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:x])+1], inorder[:x])
	root.Right = buildTree(preorder[len(inorder[:x])+1:], inorder[x+1:])
	return root
}

var preOrderRes []int

// 前序遍历: 根->左->右 递归方案
func preOrder(node *TreeNode) {
	// 节点为空
	if node != nil {
		// 访问节点数据
		preOrderRes = append(preOrderRes, node.Val)
		// 访问左子树
		preOrder(node.Left)
		// 访问右子树
		preOrder(node.Right)
	}
}

// 前序遍历 for循环方案
func preOrderByFor(node *TreeNode) {
	// 首先放根节点入栈
	Push(node)
	var cursor *TreeNode
	// 栈不为空，执行for循环
	for !IsEmptyStack() {
		cursor = GetTop()
		if cursor != nil {
			// 访问节点数据
			preOrderRes = append(preOrderRes, cursor.Val)
			// 放入左子树
			Push(cursor.Left)
		} else {
			// 将栈顶的元素移出去
			Pop()
			// pop获取栈顶元素
			cursor = Pop()
			// 将右子树进栈
			Push(cursor.Right)
		}

	}
}

var inOrderRes []int

// 中序遍历：左->根->右
func inOrder(node *TreeNode) {
	if node != nil {
		inOrder(node.Left)
		inOrderRes = append(inOrderRes, node.Val)
		inOrder(node.Right)
	}
}

var postOrderRes []int

// 后序遍历：左->右->根
func postOrder(node *TreeNode) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		postOrderRes = append(postOrderRes, node.Val)
	}
}

var preNode *ThreadTreeNode

// preOrderThread 通过前序遍历构造二叉线索树
/**
二叉树的线索化过程，分为两步
第一步：使用前序遍历，将二叉树线索化
1.定义一个全局变量，表示当前节点的前驱节点
	preNode = nil
2.使用前序遍历，递归的为每个节点的空链域设置线索
	a.给当前节点设置前驱节点
		检查当前节点的左孩子节点是否为空
		若为空，建立当前节点的左孩子节点是否为空
		若不为空，递归的访问左孩子节点
	b.给前一个节点设置后继线索
		检查前一个节点的右孩子是否为空
		若为空，建立前一个节点的右指针线索，指向当前节点
		若不为空，不做处理
	c.设当前节点为前一个节点
3.最后一个节点，单独处理右指针线索
*/
func PreOrderThread(node *ThreadTreeNode) {
	// 节点为空，什么也不做
	if node != nil {
		// a.给当前节点设置前驱节点
		if node.Left == nil {
			node.Left = preNode
			node.LTag = 1
		}
		// b.给前一个节点设置后继线索
		if preNode != nil && preNode.Right == nil {
			preNode.Right = node
			preNode.RTag = 1
		}
		// c.设置当前节点为前一个节点
		preNode = node
		// 递归访问左子树
		if node.LTag == 0 {
			PreOrderThread(node.Left)
		}
		// 递归访问右子树
		if node.RTag == 0 {
			PreOrderThread(node.Right)
		}
	}
}

// CreatePreOrderThread 构造二叉前序线索树
func CreatePreOrderThread(root *ThreadTreeNode) {
	if root != nil {
		PreOrderThread(root)
		// 最后一个节点，单独处理右指针线索
		if preNode != nil {
			preNode.RTag = 1
		}
	}

}

var preOrderThreadRes []int

// PreOrderThread 线索二叉树的前序遍历
func PreOrderThreadErgodic(node *ThreadTreeNode) {
	// 节点为空，什么都不做
	if node != nil {
		preOrderThreadRes = append(preOrderThreadRes, node.Val)
		if node.LTag == 0 {
			PreOrderThreadErgodic(node.Left)
		} else {
			PreOrderThreadErgodic(node.Right)
		}
	}
}

// PreOrderThreadByFor 循环进行遍历
func PreOrderThreadErgodicByFor(node *ThreadTreeNode) {
	cursor := node
	for cursor != nil {
		preOrderThreadRes = append(preOrderThreadRes, cursor.Val)
		if node.LTag == 0 {
			cursor = cursor.Left
		} else {
			cursor = cursor.Right
		}
	}
}

/**
中序遍历：左->根->右
二叉树的线索化过程，分为两步
第一步：使用中序序遍历，将二叉树线索化
1.定义一个全局变量，表示当前节点的前驱节点
	preNode = nil
2.使用中序遍历，递归的为每个节点的空链域设置线索
	a.给当前节点设置前驱节点
		检查当前节点的左孩子节点是否为空
		若为空，建立当前节点的左孩子节点是否为空
		若不为空，递归的访问左孩子节点
	b.给前一个节点设置后继线索
		检查前一个节点的右孩子是否为空
		若为空，建立前一个节点的右指针线索，指向当前节点
		若不为空，不做处理
	c.设当前节点为前一个节点
3.最后一个节点，单独处理右指针线索
*/
func InOrderThread(node *ThreadTreeNode) {
	// 节点为空，什么都不做
	if node != nil {
		// 递归的访问左子树
		InOrderThread(node)
		// 给当前节点设置前驱节点
		if node.Left == nil {
			node.Left = preNode
			node.LTag = 1
		}
		// 给前一个节点设置后继节点
		if preNode != nil && preNode.Right == nil {
			preNode.Right = node
			preNode.RTag = 1
		}
		// 设置当前节点为前一个节点
		preNode = node
		// 递归访问右子树
		InOrderThread(node.Right)
	}
}

// CreateInOrderThread 通过中序遍历构造二叉线索树入口函数
func CreateInOrderThread(node *ThreadTreeNode) {
	preNode = nil
	if node != nil {
		InOrderThread(node)
		if preNode != nil {
			preNode.RTag = 1
		}
	}
}

var inOrderThreadRes []int

// InOrderThreadErgodicByFor 中序遍历二叉线索树
func InOrderThreadErgodicByFor(node *ThreadTreeNode) {
	if node == nil {
		return
	}
	firstNode := node
	// 获取第一个节点
	for firstNode.LTag == 0 {
		firstNode = firstNode.Left
	}
	cursor := firstNode
	// 依次访问线索二叉树的节点
	for cursor != nil {
		// 访问节点数据
		inOrderThreadRes = append(inOrderThreadRes, cursor.Val)
		/**
		找到当前节点的后继节点
		当前节点的后继节点，即：node.Right
		由于rTag的值不同，node.Right的指向含义不同，要区别对待
		*/
		if cursor.RTag == 0 {
			// 表明有右子树
			cursor = cursor.Right
			for cursor.LTag == 0 {
				cursor = cursor.Left
			}
		} else {
			cursor = cursor.Right
		}
	}
}
