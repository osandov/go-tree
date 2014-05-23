/*
Package tree implements several tree structures.
*/
package tree

// Key in a tree.
type Key interface {
	CompareTo(Key) int
}

// Tree dynamic set.
type Tree interface {
	// Get returns the value corresponding to the given key. If the key was in
	// the tree, it returns the corresponding value and true; otherwise, it
	// returns false.
	Get(Key) (interface{}, bool)

	// Set inserts a value with the given key to a tree. If the key was already
	// in the tree, it returns the old value and true; otherwise, it returns
	// false.
	Set(Key, interface{}) (interface{}, bool)

	// Del removes the node with the given key from the tree. If the key was in
	// the tree, it returns the corresponding value and true; otherwise, it
	// returns false.
	Del(Key) (interface{}, bool)
}
