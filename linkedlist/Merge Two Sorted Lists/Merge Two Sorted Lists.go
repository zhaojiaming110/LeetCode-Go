// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/1 下午12:24

package Merge_Two_Sorted_Lists

type ListNode struct {
	Val 	int
	Next	*ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	cur := list
	value := 0
	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			value = l2.Val
			l2 = l2.Next
		} else {
			value = l1.Val
			l1 = l1.Next
		}
		node := &ListNode{Val:value}
		cur.Next = node
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return list.Next
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	}
}