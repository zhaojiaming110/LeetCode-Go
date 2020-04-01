![](https://tva1.sinaimg.cn/large/006y8mN6ly1g95r6zww73j312k0u0q77.jpg)

### 题目描述

将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

### 示例

```
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
```

### 解题

#### 迭代法

##### 解题思路

依次遍历链表`l1`和`l2`，每次遍历时比较当前节点值得大小，然后在较小值节点所在链表继续进行遍历。遍历结束后，若其中一个链表为nil，则将另一个链表赋给新链表。

##### 代码

```go
list := &ListNode{}
	cur := list
	value := 0
	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			value = l2.Val
			l2 = l2.Next
		} else {
			value = l1.Val
			l1 = l1.Next
		}
		node := &ListNode{Val:value}
		cur.Next = node
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return list.Next
```

##### 复杂度分析

- 时间复杂度：$O(n+m)$，因为每次循环迭代中，`l1`和`l2`只有一个元素会被放入合并链表中，`for`循环的次数最大值等于两个链表的总长度。所有其它工作都是常数级别的，所以总的时间复杂度是线性的。
- 空间复杂度：$O(m+n)$，新链表的长度为`m+n`。

#### 递归法

##### 解题思路

我们可以如下递归地定义在两个链表里的`merge`操作（忽略边界情况，比如空链表等）；
![](https://tva1.sinaimg.cn/large/00831rSTly1gde8cklv13j30ma029mx2.jpg)

也就是说两个链表的头部较小的一个与剩下元素的`merge`操作结果合并。

##### 代码

```go
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	}
}
```

##### 复杂度分析

- 时间复杂度：$O(n+m)$，每次调用都会去掉两个链表中头部较小的一个元素，时间复杂度与合并后的链表长度为线性关系。
- 空间复杂度：$O(n+m)$，调用 mergeTwoLists 退出时 l1 和 l2 中每个元素都一定已经被遍历过了，所以n+m个栈帧会消耗$O(n+m)$的空间。
