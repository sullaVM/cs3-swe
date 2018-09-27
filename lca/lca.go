// Package lca is the implementation for the Lowest Common Ancestor problem.
package lca

// node is a node for a binary tree.
type node struct {
	val   int
	left  *node
	right *node
}

func insert(root *node, val int) *node {
	if root == nil {
		return &node{
			val: val,
		}
	}
	if val > root.val {
		x := insert(root.right, val)
		root.right = x
	} else if val < root.val {
		x := insert(root.left, val)
		root.left = x
	} else {
		// If value is equal to root value, return root.
		return root
	}
	return nil
}

func get(root *node, val int) *node {
	if root == nil {
		return nil
	}
	if root.val == val {
		return root
	} else if root.val > val {
		return get(root.left, val)
	}
	return get(root.right, val)
}

func lowestCommonAncestor(root, p, q *node) *node {
	if root == nil || p == nil || q == nil {
		return nil
	}
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

func isAncestor(curnode, p *node) bool {
	if curnode == p {
		// A node is a descendant for itself.
		return true
	}

	leftRes := false
	if curnode.left != nil {
		if curnode.left == p {
			leftRes = true
		} else {
			leftRes = isAncestor(curnode.left, p)
		}
	}

	rightRes := false
	if curnode.right != nil {
		if curnode.right != nil {
			if curnode.right == p {
				rightRes = true
			} else {
				rightRes = isAncestor(curnode.right, p)
			}
		}
	}

	return leftRes || rightRes
}
