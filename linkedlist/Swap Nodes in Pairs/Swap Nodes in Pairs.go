// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/1 下午4:19

package Swap_Nodes_in_Pairs

type ListNode struct {
	Val 	int
	Next	*ListNode
}

// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next

	// swap node
	first.Next = swapPairs(second.Next)
	second.Next = first

	return second
}

// 迭代
func swapPairs1(head *ListNode) *ListNode {
	// 新建一个链表用来保存迭代后的结果
	dummy := &ListNode{}
	dummy.Next = head // 边界情况 若链表只有一个节点
	cur := dummy
	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		// swaping
		cur.Next = second
		first.Next, second.Next = second.Next, first

		cur = first
		head = first.Next
	}
	return dummy.Next
}