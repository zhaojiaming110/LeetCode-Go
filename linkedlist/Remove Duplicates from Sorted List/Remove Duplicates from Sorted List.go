// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/2 下午1:19

package Remove_Duplicates_from_Sorted_List

type ListNode struct {
	Val 	int
	Next	*ListNode
}

//  Input: 1->1->2
//  Output: 1->2
//	迭代法
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//	递归
// 递推公式等于，
// 当head.Val == head.Next.Val， 删除head.Next节点  head = deleteDuplicates1(head.Next)
// 当head.Val != head.Next.Val，不进行任何操作 head.Next = deleteDuplicates2(head.Next)
// [1, 1, 2, 3, 3]
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		head = deleteDuplicates1(head.Next)
	} else {
		head.Next = deleteDuplicates1(head.Next)
	}

	return head
}