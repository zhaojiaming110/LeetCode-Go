![](https://tva1.sinaimg.cn/large/006y8mN6ly1g95r6zww73j312k0u0q77.jpg)

### 题目描述

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

### 示例

```
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```

### 解题

#### 解题思路

对两个链表进行遍历，需要考虑上次遍历所产生的进位以及本次遍历可能出现的下次遍历的进位，新链表结点的值=本次遍历得到的两个结点的值 + 进位值。

#### 代码实现

```go
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
```

#### 复杂度分析

- 时间复杂度：$O(max(m,n))$，假设m和n分别为l1和l2的长度，上面的算法最多重复$max(m,n)+1$次。
- 空间复杂度：$O(max(m,n))$，新列表的长度最多为$max(m,n)+1$。




















