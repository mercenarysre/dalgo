package main

type UnionFind struct {
	Parent map[int]int
	Rank   map[int]int
}

// Constructor: initialize parent and rank
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		Parent: make(map[int]int),
		Rank:   make(map[int]int),
	}

	for i := 1; i <= n; i++ {
		uf.Parent[i] = i
		uf.Rank[i] = 0
	}

	return uf
}

// The find method accepts a vertex n as an argument and return its root parent
// not just return the direct parent
// Using the parent hashmap, the key is the vertex and the value is the parent
// If a vertex is a parent to itself, then it is the root parent
// If two vertices have the same root parent, then they are already part of
// of the same connected apartment
// Time Complexity: α(n)where we have to traverse every single node/vertex
// Space Complexity: O(1)
func (uf *UnionFind) Find(n int) int {
	p := uf.Parent[n]
	for p != uf.Parent[p] {
		uf.Parent[p] = uf.Parent[uf.Parent[p]] // path compression
		p = uf.Parent[p]
	}
	return p
}

// Union function takes two vertices and determines if a union can be performed
// if the two vertices share the same root parent, a union cannot be formed and
// we return false
// if two vertices n1 and n2 have parents p1 and p2 respectively, and p1 rank is
// higher than p2, then p2 is the child to p1
// conversely, p1 is the child to p2 if p2 rank is higher than p1
// if rank of p1 and p2 are equal, we can set p2 to the parent of p1 arbitrarily
// and increment rank by 1
// Time Complexity: α(n)where we have to traverse every single node/vertex
// Space Complexity: O(1)
func (uf *UnionFind) Union(n1, n2 int) bool {
	p1 := uf.Find(n1)
	p2 := uf.Find(n2)

	if p1 == p2 {
		return false
	}

	if uf.Rank[p1] > uf.Rank[p2] {
		uf.Parent[p2] = p1
	} else if uf.Rank[p1] < uf.Rank[p2] {
		uf.Parent[p1] = p2
	} else {
		uf.Parent[p1] = p2
		uf.Rank[p2]++
	}
	return true
}
