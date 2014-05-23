package tree

// Binary search tree implementation.

// Binary search tree.
type binarySearchTree struct {
	// Root of the tree.
	root *bstNode
}

// Node in a binary search tree.
type bstNode struct {
	key         Key
	value       interface{}
	left, right *bstNode
}

// Create a new binary search tree.
func NewBST() Tree {
	return new(binarySearchTree)
}

func (bst *binarySearchTree) Get(key Key) (interface{}, bool) {
	node := bst.root

	// Iterate down the tree.
	for {
		// We hit nil; the key isn't in the tree.
		if node == nil {
			return nil, false
		}

		cmp := node.key.CompareTo(key)
		if cmp < 0 {
			node = node.left
		} else if cmp > 0 {
			node = node.right
		} else {
			// Found it.
			return node.value, true
		}
	}
}

func (bst *binarySearchTree) Set(key Key, value interface{}) (interface{}, bool) {
	node := bst.root

	// If the root is nil, then this is the first node and therefore the new
	// root.
	if node == nil {
		bst.root = &bstNode{key, value, nil, nil}
		return nil, false
	}

	// Iterate down the tree.
	for {
		cmp := node.key.CompareTo(key)
		if cmp < 0 {
			if node.left == nil {
				node.left = &bstNode{key, value, nil, nil}
				return nil, false
			} else {
				node = node.left
			}
		} else if cmp > 0 {
			if node.right == nil {
				node.right = &bstNode{key, value, nil, nil}
				return nil, false
			} else {
				node = node.right
			}
		} else {
			orig_value := node.value
			node.key = key
			node.value = value
			return orig_value, true
		}
	}
}

func (bst *binarySearchTree) Del(key Key) (interface{}, bool) {
	var parent *bstNode
	node := bst.root

	// Iterate down the tree to find the node to remove, keeping track of the
	// parent node.
	for {
		// The key isn't in the tree.
		if node == nil {
			return nil, false
		}

		cmp := node.key.CompareTo(key)
		if cmp < 0 {
			parent, node = node, node.left
		} else if cmp > 0 {
			parent, node = node, node.right
		} else {
			// Found it.
			break
		}
	}

	var replacement *bstNode

	if node.left != nil && node.right != nil {
		// Two children.

		// Find the successor node (smallest node which is greater than the
		// target node).
		successorParent := node
		successor := node.right
		for successor.left != nil {
			successorParent, successor = successor, successor.left
		}

		// Remove it from the tree.
		if successor == successorParent.left {
			successorParent.left = successor.right
		} else {
			successorParent.right = successor.right
		}

		// Replace the node with its successor.
		successor.left = node.left
		successor.right = node.right
		replacement = successor
	} else if node.left != nil {
		// One child on the left.

		// Replace the node with its left child.
		replacement = node.left
	} else if node.right != nil {
		// One child on the right.

		// Replace the node with its right child.
		replacement = node.right
	}

	// No children and common fall-through code.

	if parent == nil {
		bst.root = replacement
	} else if node == parent.left {
		parent.left = replacement
	} else {
		parent.right = replacement
	}

	return node.value, true
}
