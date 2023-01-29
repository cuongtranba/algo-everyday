package dailycodingproblem

//Given a linked list,
//rearrange the node values such that they appear in alternating low -> high -> low -> high ... form.
//For example, given 1 -> 2 -> 3 -> 4 -> 5, you should return 1 -> 3 -> 2 -> 5 -> 4.
type Node struct {
	Val  int
	Next *Node
}

func SortLowAndHigh(n *Node) {
	if n == nil {
		return
	}
	count := 0
	head := n
	for n.Next != nil {
		if count%2 != 0 {
			n.Val, n.Next.Val = n.Next.Val, n.Val
		}
		count++
		n = n.Next
	}
	n = head
}
