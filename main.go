package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// BottomView returns, for every vertical column (root at 0, left children
// -1, right +1), the value seen from below — the LAST node BFS meets in that
// column, since BFS visits top to bottom.
func BottomView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	type item struct {
		n  *TreeNode
		hd int // horizontal distance from the root column
	}
	last := map[int]int{} // column -> value of the deepest (latest in BFS) node
	minHd, maxHd := 0, 0
	queue := []item{{root, 0}}
	for len(queue) > 0 {
		it := queue[0]
		queue = queue[1:]
		last[it.hd] = it.n.Val // later rows overwrite earlier ones
		if it.hd < minHd {
			minHd = it.hd
		}
		if it.hd > maxHd {
			maxHd = it.hd
		}
		if it.n.Left != nil {
			queue = append(queue, item{it.n.Left, it.hd - 1})
		}
		if it.n.Right != nil {
			queue = append(queue, item{it.n.Right, it.hd + 1})
		}
	}
	out := make([]int, 0, maxHd-minHd+1)
	for hd := minHd; hd <= maxHd; hd++ {
		out = append(out, last[hd])
	}
	return out
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(20, tn(8, tn(5, nil, nil), tn(3, tn(10, nil, nil), tn(14, nil, nil))), tn(22, nil, tn(25, nil, nil)))
	fmt.Println(BottomView(root))
}
