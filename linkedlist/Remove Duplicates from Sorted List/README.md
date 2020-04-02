
![](https://tva1.sinaimg.cn/large/006y8mN6ly1g95r6zww73j312k0u0q77.jpg)

## 题目描述

给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

## 示例

```
输入: 1->1->2->3->3
输出: 1->2->3
```

## 解题

### 迭代法

#### 解题思路

直接遍历链表，比较当前节点和下一节点的值即可。

#### 代码

```go
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
```

#### 复杂度分析

- 时间复杂度：O(N)，链表的一趟遍历。
- 空间复杂度：O(1)。

### 递归法

#### 解题思路

我们可以如下定义递归在函数`deleteDuplicates`里的操作。

递推公式分如下两种情况：

1. 当`head.Val == head.Next.Val`，删除`head.Next`节点  +  `head = deleteDuplicates(head.Next)`
2. 当`head.Val != head.Next.Val`，不进行任何操作 +  `head.Next = deleteDuplicates2(head.Next)`

#### 代码

```go
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		head = deleteDuplicates(head.Next)
	} else {
		head.Next = deleteDuplicates(head.Next)
	}

	return head
}
```

#### 复杂度分析

- 时间复杂度：O(N)，递归次数等于链表的长度。
- 空间复杂度：O(N)，递归过程中使用的堆栈空间。


