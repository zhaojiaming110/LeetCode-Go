![](https://tva1.sinaimg.cn/large/006y8mN6ly1g95r6zww73j312k0u0q77.jpg)
### 题目描述

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

**你不能只是单纯的改变节点内部的值**，而是需要实际的进行节点交换。

### 示例

```
给定 1->2->3->4, 你应该返回 2->1->4->3.
```

### 解题

#### 递归法

##### 解题思路

我们可以如下定义递归在`swapPairs`里的操作（忽略边界情况，如`head==nil || head.Next == nil`）
$
递推公式 = secondNode + firstNode + swapPairs(secondNode.Next)
$

##### 代码

```go
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next

	// swap node
	first.Next = swapPairs(second.Next)
	second.Next = first

	return second
}
```

##### 复杂度分析

- 时间复杂度：O(N)，其中N指的是链表的节点数量。
- 空间复杂度：O(N)，递归过程使用的堆栈空间。

#### 迭代法

##### 解题思路

新建一个链表用来保存迭代后的结果，用`cur`表示当前节点

在每次迭代过程中，先使`cur.Next = secondNode`，交换`firstNode`和`secondNode`节点的位置。

交换完成后，为下一次交换进行初始化操作，令`cur = firstNode`，`head = firstNode.Next`

注意边界条件，若链表只有一个节点的情况。

##### 代码

```go
func swapPairs(head *ListNode) *ListNode {
	// 新建一个链表用来保存迭代后的结果
	dummy := &ListNode{}
	dummy.Next = head // 边界情况 若链表只有一个节点
	cur := dummy
	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		// swaping
		cur.Next = second
		first.Next, second.Next = second.Next, first

		cur = first
		head = first.Next
	}
	return dummy.Next
}
```

##### 复杂度分析

- 时间复杂度：O(N)，N为链表的节点个数。
- 空间复杂度：O(1)，我们只使用了常量级的额外空间。


