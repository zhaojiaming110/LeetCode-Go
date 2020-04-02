// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/2 下午2:45

package Partition_List

type ListNode struct {
	Val 	int
	Next	*ListNode
}

func partition(head *ListNode, x int) *ListNode {
	front_head := &ListNode{}
	after_head := &ListNode{}
	front := front_head
	after := after_head

	// 遍历链表 生成两个子链表 front_head.Next和after_head.Next
	for head != nil {
		if head.Val < x {
			front.Next = head
			front = front.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	// 节点after可能并不是原链表的最后一个节点
	after.Next = nil
	// 拼接链表
	front.Next = after_head.Next

	return front_head.Next
}
