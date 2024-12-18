package main

type ListDeleter struct {}

// 删除排序链表中的重复元素
func (deleter *ListDeleter) DeleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

// 删除排序链表中的重复元素 II
func (deleter *ListDeleter) DeleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	current := dummy
	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			value := current.Next.Val
			for current.Next != nil && current.Next.Val == value {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}
	}
	return dummy.Next
}

// 删除链表的倒数第 N 个节点: 计算长度
func (deleter *ListDeleter) RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	length := deleter.getLength(head)
	dummy := &ListNode{Val: 0, Next: head}
	current := dummy
	for index := 0; index < length - n; index++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	return dummy.Next
}

// 删除链表的倒数第 N 个节点: 双指针
func (deleter *ListDeleter) RemoveNthFromEndDoublePointer(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	first, second := head, dummy
	for index := 0; index < n; index++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func (deleter *ListDeleter) getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}
