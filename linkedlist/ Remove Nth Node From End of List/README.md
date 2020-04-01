![](https://tva1.sinaimg.cn/large/006y8mN6ly1g95r6zww73j312k0u0q77.jpg)

### 题目描述

给定一个链表，删除链表的倒数第 *n* 个节点，并且返回链表的头结点。

### 示例

```
给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
```

### 解题

#### 两次遍历

##### 解题思路

删除链表的倒数第n个节点，可以转化为删除链表的第$L-n+1$个节点。

预先设置一个哑结点，哑结点用来简化要删除头结点或链表只有一个节点的情况。

##### 解题

```go
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
```

##### 复杂度分析

- 时间复杂度：$O(L)$，该算法对链表进行了两次遍历，其中计算链表的长度时遍历了L次，其次找到要删除的前节点遍历了L-n次，操作执行了2L-n次，时间复杂度为$O(L)$。
- 空间复杂度：$O(1)$，我们只用了常量级的额外空间。

#### 一次遍历

##### 解题思路

- 预先设置一个哑结点
- 定义`first`指针和`last`指针，其中`first`指针遍历后`first=nil`，`last`指针遍历后指向要删除节点的前一个节点
- 根据最后两个指针的位置可以得出，`first`指针需要先走`n+1`步。

##### 解题

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
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
```

##### 复杂度分析

- 时间复杂度：$O(L)$，该算法对含有L个节点的链表进行了一次遍历。
- 空间复杂度：$O(1)$，我们只用了常量级的额外空间。