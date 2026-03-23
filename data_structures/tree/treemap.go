package main

type MapNode struct {
	Key   string
	Value int
	Left  *MapNode
	Right *MapNode
}

type TreeMap struct {
	Root *MapNode
}

// Time Complexity: O(logn)
// Space Complexity: O(h)
func (root *MapNode) Insert(key string, value int) *MapNode {
	if root == nil {
		return &MapNode{
			Key:   key,
			Value: value,
			Left:  nil,
			Right: nil,
		}
	}

	if key < root.Key {
		root.Left = root.Left.Insert(key, value)
	} else if key > root.Key {
		root.Right = root.Right.Insert(key, value)
	} else {
		root.Value = value
	}
	return root
}

// Time Complexity: O(logn)
// Space Complexity: O(h)
func (root *MapNode) Search(key string) (int, bool) {
	if root == nil {
		return -1, false
	}

	if key < root.Key {
		return root.Left.Search(key)
	} else if key > root.Key {
		return root.Right.Search(key)
	}
	return root.Value, true
}

// Time Complexity: O(logn)
// Space Complexity: O(h)
func (t *TreeMap) Insert(key string, value int) {
	t.Root = t.Root.Insert(key, value)
}

// Time Complexity: O(logn)
// Space Complexity: O(h)
func (t *TreeMap) Search(key string) (int, bool) {
	return t.Root.Search(key)
}
