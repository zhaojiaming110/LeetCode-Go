
![](https://tva1.sinaimg.cn/large/006y8mN6ly1g95r6zww73j312k0u0q77.jpg)

## 题目描述

给定一个链表，旋转链表，将链表每个节点向右移动 *k* 个位置，其中 *k* 是非负数。

## 示例

```
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
```

## 解题

### 解题思路

首先考虑如何找出旋转位置，剩下的只需将链表首尾相连形成闭环，在旋转位置将闭环链表分成两部分。

### 解题

```go
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
```

### 复杂度分析

- 时间复杂度：O(N)，链表共有N个节点。
- 空间复杂度：O(1)，只需要常量级的空间。


