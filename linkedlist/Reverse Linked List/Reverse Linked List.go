// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/2 下午3:39

package Reverse_Linked_List

type ListNode struct {
	Val 	int
	Next	*ListNode
}

// 迭代
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
func reverseList(head *ListNode) *ListNode {
	var preNode *ListNode
	for head != nil {
	 	nextNode := head.Next
		head.Next = preNode
		preNode = head
		head = nextNode
	}

	return preNode
}

// 递归
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nextNode := reverseList1(head.Next)
	head.Next.Next = head	// 下一节点的Next指向当前节点
	head.Next = nil	// 使其不指向原链表的下一个节点 当前节点的Next在上一层递归中指向

	return nextNode
}