package main

type SegmentTree struct {
	Sum          int
	LeftPointer  *SegmentTree
	RightPointer *SegmentTree
	L, R         int
}

// Build constructs the segment tree
// Time Complexity: O(n): n denotes number of nodes the tree contains
func Build(nums []int, L, R int) *SegmentTree {
	if L == R {
		return &SegmentTree{
			Sum: nums[L],
			L:   L,
			R:   R,
		}
	}

	M := (L + R) / 2
	root := &SegmentTree{
		L: L,
		R: R,
	}
	root.LeftPointer = Build(nums, L, M)
	root.RightPointer = Build(nums, M+1, R)

	// rangeSum of the SegmentTree from bottom up
	root.Sum = root.LeftPointer.Sum + root.RightPointer.Sum
	return root
}

// Update updates the value at index to val
// Time Complexity: O(log n), Going down the height of the tree
// a balanced tree, h is height of the tree, known as logn
// Update, a recursive function, takes in the index and the value
// we want to update it with, our base case is if we can reach a leaf node
// we have a found our index and we can update the sum to the new value
// Then we call the M value; M < index == go to right subtree, M > index == go to left subtree
func (t *SegmentTree) Update(index, val int) {
	if t.L == t.R {
		t.Sum = val
		return
	}

	M := (t.L + t.R) / 2
	if index > M {
		t.RightPointer.Update(index, val)
	} else {
		t.LeftPointer.Update(index, val)
	}
	t.Sum = t.LeftPointer.Sum + t.RightPointer.Sum
}

// RangeQuery queries the sum in range [L, R]
// Time Complexity: O(log n), Going down the height of the tree
// a balanced tree, h is height of the tree, known as logn
// Following a recursive procedure, we recurse down the tree
// and calculate M (mid value), Given L and R, if L > M, range lies
// on the right, if L < M, range lies on the left
// Base Case
// Range is in the right subtree
// Range is in the left subtree
// Range overlaps with the right and left subtrees
func (t *SegmentTree) RangeQuery(L, R int) int {
	if L == t.L && R == t.R {
		return t.Sum
	}

	M := (t.L + t.R) / 2
	if L > M {
		return t.RightPointer.RangeQuery(L, R)
	} else if R <= M {
		return t.LeftPointer.RangeQuery(L, R)
	} else {
		return t.LeftPointer.RangeQuery(L, M) +
			t.RightPointer.RangeQuery(M+1, R)
	}
}
