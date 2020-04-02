// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/1 下午6:33

package _Rotate_List

type ListNode struct {
	Val 	int
	Next	*ListNode
}

// 1->2->3->4->5->6NULL, k = 2
// 5->6->1->2->3->4->NULL
func rotateRight(head *ListNode, k int) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}

	// 首先计算链表的长度
	length := 1
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
		length++
	}
	// 使链表首尾相连，形成闭环
	tmp.Next = head

	// 重新计算k的大小
	k %= length
	// 下面注释的代码在leetCode会超时，应该是LeetCode的bug.
	//if k == 0 {
	//	return head
	//}

	// 找出旋转位置 将初始链表分成两部分
	cur := head
	for i := 0; i < length - k - 1; i++ {
		cur = cur.Next
	}

	result := cur.Next
	cur.Next = nil

	return result
}