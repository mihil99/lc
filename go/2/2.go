package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t := l1.Val + l2.Val
	r := &ListNode{t % 10, nil}
	cur := r
	for l1.Next != nil || l2.Next != nil || t/10 != 0 {
		t = t / 10
		if l1.Next != nil {
			l1 = l1.Next
			t += l1.Val
		}
		if l2.Next != nil {
			l2 = l2.Next
			t += l2.Val
		}
		cur.Next = &ListNode{t % 10, nil}
		cur = cur.Next
	}
	return r
}

func main() {
	fmt.Println(addTwoNumbers(createListNode([]int{2, 4, 3}), createListNode([]int{5, 6, 4})).String())
	fmt.Println(addTwoNumbers(createListNode([]int{0}), createListNode([]int{0})).String())
	fmt.Println(addTwoNumbers(createListNode([]int{9, 9, 9, 9, 9, 9, 9}), createListNode([]int{9, 9, 9, 9})).String())
}

func (listNode *ListNode) String() string {
	curListNode := listNode
	s := fmt.Sprintf("[%d", curListNode.Val)
	for curListNode.Next != nil {
		curListNode = curListNode.Next
		s += fmt.Sprintf(" ,%d", curListNode.Val)
	}
	s += "]\n"
	return s
}

func createListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	return &ListNode{nums[0], createListNode(nums[1:])}
}
