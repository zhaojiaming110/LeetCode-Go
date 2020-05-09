// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/2 下午5:23

package Reverse_Linked_List_II

type ListNode struct {
	Val 	int
	Next	*ListNode
}

// 迭代
// Input: 1->2->3->4->5->NULL, m = 2, n = 4
// Output: 1->4->3->2->5->NULL
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// base case
	if head == nil || head.Next == nil || m == n {
		return head
	}

	var preNode *ListNode
	preNode = head
	var node1 *ListNode
	var node2 *ListNode
	var node3 *ListNode
	var node4 *ListNode
	cur := head
	for i := 1; cur != nil; i++ {
		if i < m {
			preNode = cur
			cur = cur.Next
			continue
		}
		if i == m {
			node1 = preNode
			node2 = cur
			preNode = cur
			cur = cur.Next
			continue
		}
		if i > n {
			break
		}
		if i == n {
			node3 = cur
			node4 = cur.Next
			nextNode := cur.Next
			cur.Next = preNode
			cur = nextNode
			continue
		}
		nextNode := cur.Next
		cur.Next = preNode
		preNode = cur
		cur = nextNode
	}

	if m == 1 {
		node1.Next = node3
	} else {
		node1.Next = node3
		node2.Next = node4
	}

	return head
}