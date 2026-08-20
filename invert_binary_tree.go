package algoritmosexercicios

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	q := make([]*TreeNode, 0)

	q = append(q, root)

	for len(q) > 0 {
		//dequeue
		node := q[0]
		q = q[1:]
		if node != nil {
			node.Left, node.Right = node.Right, node.Left
			//enqueue
			q = append(q, node.Left, node.Right)
		}

	}
	return root
}

//recursive

func invertTreeRec(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTreeRec(root.Right), invertTreeRec(root.Left)

	return root
}
