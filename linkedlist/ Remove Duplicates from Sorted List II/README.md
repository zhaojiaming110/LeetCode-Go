![](https://tva1.sinaimg.cn/large/006y8mN6ly1g95r6zww73j312k0u0q77.jpg)
## 题目描述

给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 *没有重复出现* 的数字。

## 示例

```
输入: 1->2->3->3->4->4->5
输出: 1->2->5
```

## 解题

### 迭代法

#### 解题思路

使用快慢指针，用快指针跳过哪些有重复的数组，慢指针负责和快指针进行拼接。

在链表前新插入一个哑结点，哑结点用来简化要删除头结点的这种情况。同时注意删除尾结点的情况。

#### 解题

```go
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
```

#### 复杂度分析

- 时间复杂度：O(N)，链表结点个数为N。
- 空间复杂度：O(1)，只使用了常量级的空间。

### 递归法

#### 解题思路

我们可以如下定义递归在`deleteDuplicates`里的操作，令`fastNode = head.Next`

递推公式分以下两种情况

1. 当`fastNode.Val != head.Val`，`head + deleteDuplicates(head.Next)`；
2. 当`fastNode.Val == head.Val`，`head + [删除重复元素] + deleteDuplicates(fastNode)`；

#### 解题

```go
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fastNode := head.Next
	if fastNode.Val == head.Val {
		for fastNode != nil && head.Val == fastNode.Val {
			fastNode = fastNode.Next
		}
		head = deleteDuplicates(fastNode)
	} else {
		head.Next = deleteDuplicates(head.Next)
	}

	return head
}
```

#### 复杂度分析

- 时间复杂度：O(N)，递归次数与链表的长度成线性关系。
- 空间复杂度：O(N)，递归过程使用的堆栈空间。


