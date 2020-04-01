// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/1 上午11:07

package _Remove_Nth_Node_From_End_of_List

type ListNode struct {
	Val 	int
	Next	*ListNode
}

// 1 2 3 4 5  n=2
// 1 2 3   5
// - 1 2 3 4 5
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 预先设置一个哑结点
	dummy := &ListNode{Next:head}

	// 计算链表的长度
	length := 0
	for head != nil {
		length++
		head = head.Next
	}

	// 找出要删除的节点的前一个节点
	foundNode := dummy
	for i := 0; i < length - n; i++ {
		foundNode = foundNode.Next
	}
	foundNode.Next = foundNode.Next.Next

	return dummy.Next
}

// - 2 3 1 4 5   n=2
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// 预先设置一个哑结点
	dummy := &ListNode{Next:head}
	first := dummy
	last := dummy

	// 第一个指针需要先走n+1步
	for i := 0; i < n + 1; i++ {
		first = first.Next
	}

	// 两个指针同时移动，直到第一个指针为nil，此时末指针刚好指向要删除节点的前一个节点
	for first != nil {
		first = first.Next
		last = last.Next
	}

	last.Next = last.Next.Next

	return dummy.Next
}