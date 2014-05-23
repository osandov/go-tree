package tree

// Binary search tree implementation.

// Binary search tree.
type binarySearchTree struct {
	// Root of the tree.
	root *bstNode
}

// Node in a binary search tree.
type bstNode struct {
	key                 Key
	value               interface{}
	parent, left, right *bstNode
}

// Create a new binary search tree.
func NewBST() Tree {
	return new(binarySearchTree)
}

func (bst *binarySearchTree) Get(key Key) (interface{}, bool) {
	node := bst.get(key)

	if node == nil {
		return nil, false
	} else {
		return node.value, true
	}
}

// get finds the node with the given key.
func (bst *binarySearchTree) get(key Key) *bstNode {
	node := bst.root

	// Iterate down the tree.
	for {
		// We hit nil; the key isn't in the tree.
		if node == nil {
			return nil
		}

		cmp := node.key.CompareTo(key)
		if cmp < 0 {
			node = node.left
		} else if cmp > 0 {
			node = node.right
		} else {
			// Found it.
			return node
		}
	}
}

func (bst *binarySearchTree) Set(key Key, value interface{}) (interface{}, bool) {
	node, exists := bst.add(key)

	if exists {
		orig_value := node.value
		node.value = value
		return orig_value, true
	} else {
		node.value = value
		return nil, false
	}
}

// add returns the node containing the given key or creates one and returns it.
func (bst *binarySearchTree) add(key Key) (*bstNode, bool) {
	node := bst.root

	// If the root is nil, then this is the first node and therefore the new
	// root.
	if node == nil {
		bst.root = &bstNode{key, nil, nil, nil, nil}
		return bst.root, false
	}

	// Iterate down the tree.
	for {
		cmp := node.key.CompareTo(key)
		if cmp < 0 {
			if node.left == nil {
				node.left = &bstNode{key, nil, node, nil, nil}
				return node.left, false
			} else {
				node = node.left
			}
		} else if cmp > 0 {
			if node.right == nil {
				node.right = &bstNode{key, nil, node, nil, nil}
				return node.right, false
			} else {
				node = node.right
			}
		} else {
			node.key = key
			return node, true
		}
	}
}

func (bst *binarySearchTree) Del(key Key) (interface{}, bool) {
	node := bst.del(key)

	if node == nil {
		return nil, false
	} else {
		return node.value, true
	}
}

// del removes the node with the given key and returns it.
func (bst *binarySearchTree) del(key Key) *bstNode {
	node := bst.get(key)
	if node == nil {
		return nil
	}

	var replacement *bstNode

	if node.left != nil && node.right != nil {
		// Two children.

		// Find the successor node (smallest node which is greater than the
		// target node).
		successor := node.right
		for successor.left != nil {
			successor = successor.left
		}

		// Remove it from the tree.
		if successor == successor.parent.left {
			successor.parent.left = successor.right
		} else {
			successor.parent.right = successor.right
		}
		if successor.right != nil {
			successor.right.parent = successor.parent
		}

		// Adopt the node's children.
		successor.left = node.left
		if successor.left != nil {
			successor.left.parent = successor
		}
		successor.right = node.right
		if successor.right != nil {
			successor.right.parent = successor
		}

		// Replace the node with its successor.
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

	if node.parent == nil {
		bst.root = replacement
	} else if node == node.parent.left {
		node.parent.left = replacement
	} else {
		node.parent.right = replacement
	}
	if replacement != nil {
		replacement.parent = node.parent
	}

	return node
}
