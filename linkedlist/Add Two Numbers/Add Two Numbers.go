// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/3/31 下午5:33

package Add_Two_Numbers

type ListNode struct {
	Val 	int
	Next	*ListNode
}
//2 4 3
//5 6 4
//7 0 8
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	var cur	= list
	tag := 0
	for l1 != nil || l2 != nil || tag != 0 {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b
		if tag == 1 {
			sum = sum + 1
			tag = tag - 1
		}
		if sum >= 10 {
			sum = sum - 10
			tag = tag + 1
		}
		cur.Next = &ListNode{Val:sum}
		cur = cur.Next
	}

	return list.Next
}