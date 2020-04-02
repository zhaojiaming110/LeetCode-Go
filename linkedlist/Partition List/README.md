![](https://tva1.sinaimg.cn/large/006y8mN6ly1g95r6zww73j312k0u0q77.jpg)

## 题目描述

给定一个链表和一个特定值 *x*，对链表进行分隔，使得所有小于 *x* 的节点都在大于或等于 *x* 的节点之前。

你应当保留两个分区中每个节点的`初始相对位置`。

## 示例

```
输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
```

## 解题

### 双指针法

#### 解题思路

本题要求改变链表原有数据的链表结构，使得值小于`x`的元素，位于值大于等于`x`元素的前面。并且原来链表中每个节点的`初始相对位置`保持不变。

这实质上意味着，改变后的链表结构必然存在某个`点`，在该点之前的节点的值全部小于`x`，在该点之后的元素全部大于等于`x`。

对该问题采用逆向工程，可以得出，如果我们在这个`点`将链表拆分，我们可以得到两个更小的链表，其中前半部分的链表中节点的值都小于`x`，后半部分链表中节点的值都大于等于`x`。

因此，我们只需要创建两个链表并遍历原链表，将小于`x`的节点依次添加到前链表中，大于等于`x`的节点依次添加到后链表中去，然后连接两个链表就可得到问题的解。

#### 代码

```go
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
```

#### 复杂度分析

- 时间复杂度：O(N)，遍历一趟原链表。
- 空间复杂度：O(1)，没有申请任何新的空间。在创建两个子链表的过程中，我们只是移动了原有的节点，因此没有使用任何额外的空间。