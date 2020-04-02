// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/2 上午10:17

package _Remove_Duplicates_from_Sorted_List_II

type ListNode struct {
	Val 	int
	Next	*ListNode
}

// 双指针
// Input: +->1->2->3->3->4->4->5
// Output: 1->2->5
// Input: 1->1->1->2->3
// Output: 2->3
// +->1->1->nil
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 定义哑结点
	dummy := &ListNode{Next:head}
	// 定义快慢指针
	quick := head.Next
	slow := dummy
	tag := false
	for quick != nil {
		if quick.Val == slow.Next.Val {
			// 快指针后移一位，慢指针不变
			quick = quick.Next
			tag = true
		} else {
			if tag == true {
				// 进行删除操作
				slow.Next = quick
				// 快指针向后移动一位，慢指针不动
				quick = quick.Next
				tag = false
			} else {
				// 快慢指针同时向后移动一位
				quick = quick.Next
				slow = slow.Next
			}
		}
	}
	if tag == true {
		slow.Next = quick
	}

	return dummy.Next
}



// 递归
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fastNode := head.Next
	if fastNode.Val == head.Val {
		for fastNode != nil && head.Val == fastNode.Val {
			fastNode = fastNode.Next
		}
		head = deleteDuplicates1(fastNode)
	} else {
		head.Next = deleteDuplicates1(head.Next)
	}

	return head
}