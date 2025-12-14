package main

// Graph is a data structure that is made up of nodes(like linked list and binary trees) whereby
// they are connected by edges; data structure that consists of connected nodes
// Directed graphs, Undirected graphs
type GraphNode struct { // Vertex
	Val int // Vertex value
	// other vertex properties
	Neighbors []*GraphNode // Vertex array of adjacent vertices/neighbours
}

func NewGraphNode(val int) *GraphNode {
	return &GraphNode{
		Val:       val,
		Neighbors: make([]*GraphNode, 0),
	}
}

// Matrix: storing a vertex adjacent vertices, Space Complexity: O(n * m) n: no of rows, m: no of columns
// Adjacency Matrix: storing a vertex adjacent vertices, Space Complexity: O(v>2), v is number of vertices
type Intro struct {
	// Matrix (2D Grid)
	Grid [][]int

	// Adjacency matrix
	AdjMatrix [][]int
}

func NewIntro() *Intro {
	return &Intro{
		Grid: [][]int{
			{0, 0, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 1},
			{0, 1, 0, 0},
		},
		AdjMatrix: [][]int{
			{0, 0, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 1},
			{0, 1, 0, 0},
		},
	}
}
