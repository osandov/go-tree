/*
Package tree implements several tree structures.
*/
package tree

type Key interface {
	CompareTo(Key) int
}

type Tree interface {
	// Get the value corresponding to the given key. If the key was in the
	// tree, return the corresponding value and true; otherwise, return false.
	Get(Key) (interface{}, bool)

	// Insert a value with the given tree to a key. If the key was already in
	// the tree, return the old value and true; otherwise, return false.
	Set(Key, interface{}) (interface{}, bool)

	// Remove the node with the given key from the tree. If the key was in the
	// tree, return the corresponding value and true; otherwise, return false.
	Del(Key) (interface{}, bool)
}
