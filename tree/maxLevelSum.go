package tree

import "math"

// Input: root = [1,7,0,7,-8,null,null]
// Output: 2
// Explanation:
// Level 1 sum = 1.
// Level 2 sum = 7 + 0 = 7.
// Level 3 sum = 7 + -8 = -1.
// So we return the level with the maximum sum which is level 2.

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLevel := 1
	currentLevel := 1
	maxSum := math.MinInt
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val
			if node.Left != nil {
				// push into queue to calculate next level
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				// push into queue to calculate next level
				queue = append(queue, node.Right)
			}
		}
		// Cập nhật level có tổng lớn nhất
		if levelSum > maxSum {
			maxSum = levelSum
			maxLevel = currentLevel
		}
		currentLevel++ // Chuyển sang level tiếp theo
	}

	return maxLevel
}
