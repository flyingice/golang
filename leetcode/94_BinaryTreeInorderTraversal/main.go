package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalHelper(root *TreeNode, out chan<- int) {
	if root == nil {
		close(out)
		return
	}

	left, right := make(chan int), make(chan int)
	go inorderTraversalHelper(root.Left, left)
	go inorderTraversalHelper(root.Right, right)

	for x := range left {
		out <- x
	}
	out <- root.Val
	for x := range right {
		out <- x
	}
	close(out)
}

func inorderTraversal(root *TreeNode) []int {
	ch := make(chan int)
	go inorderTraversalHelper(root, ch)

	var res []int
	for x := range ch {
		res = append(res, x)
	}

	return res
}

func main() {
}
