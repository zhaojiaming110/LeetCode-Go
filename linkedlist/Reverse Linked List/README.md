![](https://tva1.sinaimg.cn/large/006y8mN6ly1g95r6zww73j312k0u0q77.jpg)

## 题目描述

反转一个单链表，你可以迭代或递归地反转单链表。

## 示例

```
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
```

## 解题

### 迭代法

#### 解题思路

遍历链表，将当前节点的`Next`指针指向前一个节点。由于是单链表，所以事先需存储节点前的一个元素（头结点前的节点为`nil`）。在更改`Next`之前还需要保存下一个节点。

#### 代码

```go
func reverseList(head *ListNode) *ListNode {
	var preNode *ListNode
	for head != nil {
	 	nextNode := head.Next
		head.Next = preNode
		preNode = head
		head = nextNode
	}

	return preNode
}
```

#### 复杂度分析

- 时间复杂度：O(N)，遍历次数为链表的长度。
- 空间复杂度：O(1)。

### 递归

#### 解题思路

递归版本其关键在于反向工作，假设链表为：

​	`n1->...->n(k-1)->n(k)->n(k+1)...->n(m)->nil`

若从节点`n(k+1)`到`n(m)`已经被反转，我们正处于节点`n(k)`。

​	`n1->...->n(k-1)->n(k)->n(k+1)...<-n(m)->nil`

我们希望`n(k+1)`的`Next`指向`n(k)`。

所以，`n(k).Next.Next = n(k)`。也就是说当前节点`Next`在上一层递归中被修改。

为了使得`n1->nil`，在当前节点应`n(k).Next = nil`。（实际`Next`修改在上一层递归中）

此时链表为：

​	`n1->...->n(k-1)    n(k)<-n(k+1)...<-n(m)->nil`

注意这会儿链表已经断开了。

#### 代码

```go
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nextNode := reverseList(head.Next)
	head.Next.Next = head	// 下一节点的Next指向当前节点
	head.Next = nil	// 使其不指向原链表的下一个节点 当前节点的Next在上一层递归中指向

	return nextNode
}
```

#### 复杂度分析

- 时间复杂度：O(N)，递归的次数为链表的长度。
- 空间复杂度：O(N)，递归过程中使用的堆栈空间。






