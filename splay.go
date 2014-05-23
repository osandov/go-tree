package tree

// Splay tree implementation.

// Splay tree.
type splayTree struct {
	// Underlying binary search tree.
	bst binarySearchTree
}

// NewSplay creates an empty splay tree. A splay tree is a self-adjusting
// variant of a binary search tree that optimizes for locality of reference. It
// has amortized O(log n) behavior in the worst case.
func NewSplay() Tree {
	return new(splayTree)
}

func (s *splayTree) Get(key Key) (interface{}, bool) {
	node := s.bst.get(key)

	if node == nil {
		return nil, false
	} else {
		s.splayNode(node)
		return node.value, true
	}
}

func (s *splayTree) Set(key Key, value interface{}) (interface{}, bool) {
	node, exists := s.bst.add(key)

	s.splayNode(node)

	if exists {
		orig_value := node.value
		node.value = value
		return orig_value, true
	} else {
		node.value = value
		return nil, false
	}
}

func (s *splayTree) Del(key Key) (interface{}, bool) {
	node := s.bst.del(key)

	if node == nil {
		return nil, false
	} else {
		if node.parent != nil {
			s.splayNode(node.parent)
		}
		return node.value, true
	}
}

// splayNode moves a node to the root of a tree in a manner that keeps recently
// splayed elements near the root.
func (s *splayTree) splayNode(node *bstNode) {
	// Carry out splay steps until the node reaches the root.
	for node != s.bst.root {
		var parent, grandparent *bstNode
		parent = node.parent
		if parent != nil {
			grandparent = parent.parent
		}

		switch {
		// Zig step.
		case parent == s.bst.root && node == parent.left:
			s.bst.rotateRight(parent)
		case parent == s.bst.root && node == parent.right:
			s.bst.rotateLeft(parent)
		// Zig-zig step.
		case node == parent.left && parent == grandparent.left:
			s.bst.rotateRight(grandparent)
			s.bst.rotateRight(parent)
		case node == parent.right && parent == grandparent.right:
			s.bst.rotateLeft(grandparent)
			s.bst.rotateLeft(parent)
		// Zig-zag step.
		case node == parent.left && parent == grandparent.right:
			s.bst.rotateRight(parent)
			s.bst.rotateLeft(grandparent)
		case node == parent.right && parent == grandparent.left:
			s.bst.rotateLeft(parent)
			s.bst.rotateRight(grandparent)
		}
	}
}
