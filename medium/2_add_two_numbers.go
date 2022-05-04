package medium

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	curr := result
	add, sum := 0, 0
	var l, r int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			l = l1.Val
		} else {
			l = 0
		}
		if l2 != nil {
			r = l2.Val
		} else {
			r = 0
		}
		sum = l + r + add
		curr.Next = &ListNode{Val: sum % 10}
		add = sum / 10
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if add != 0 {
		curr.Next = &ListNode{Val: add}
	}
	return result.Next
}
