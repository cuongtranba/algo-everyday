package dailycodingproblem

import (
	"fmt"
	"testing"
)

func TestSortLowAndHigh(t *testing.T) {
	n := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
				Next: &Node{
					Val: 4,
					Next: &Node{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	SortLowAndHigh(n)
	fmt.Println(n)
}
