// lca is the implementation for the Lowest Common Ancestor problem.
package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func lowestCommonAncestor(root, p, q *node) *node {
	if root.left == nil && root.right == nil {
		if isAncestor(root, p) && isAncestor(root, q) {
			return root
		}
		return nil
	}

	if root.left != nil {
		if isAncestor(root.left, p) && isAncestor(root.left, q) {
			return lowestCommonAncestor(root.left, p, q)
		}
	}

	if isAncestor(root, p) && isAncestor(root, q) {
		return root
	}

	return nil
}

func isAncestor(curNode, p *node) bool {
	if curNode == p {
		// A node is a descendant for itself.
		return true
	}

	leftRes := false
	if curNode.left != nil {
		if curNode.left == p {
			leftRes = true
		} else {
			leftRes = isAncestor(curNode.left, p)
		}
	}

	rightRes := false
	if curNode.right != nil {
		if curNode.right != nil {
			if curNode.right == p {
				rightRes = true
			} else {
				rightRes = isAncestor(curNode.right, p)
			}
		}
	}

	return leftRes || rightRes
}

func main() {
	fmt.Printf("Hello world")
}
