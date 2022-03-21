package twopointer

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func countNode(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

func middleNode(head *ListNode) *ListNode {
	totalNode := countNode(head)
	for i := 0; i < totalNode/2; i++ {
		head = head.Next
	}

	return head
}
