package main

import (
	"container/list"
)

// GraphNode used for adjacency list
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func NewGraphNode(val int) *GraphNode {
	return &GraphNode{
		Val:       val,
		Neighbors: make([]*GraphNode, 0),
	}
}

// Adjacency List: storing a vertex adjacent vertices
// Space Complexity: O(v + e), v: no of vertices, e: no of edges
type AdjacencyList struct{}

// constructor function
func NewAdjacencyList() *AdjacencyList {
	return &AdjacencyList{}
}

// Using an Hashmap to implement a Graph(which has a vertex and its adjacency list)
// A hashmap where the key is a vertex and it maps to a list or array of its neighbors, which are also vertices.
// A hash map works here because we are assuming that all of the values keys are unique.
// Given directed/undirected edges, build an adjacency list
func (al *AdjacencyList) BuildAdjList() map[string][]string {
	adjList := make(map[string][]string)
	edges := [][]string{{"A", "B"}, {"B", "C"}, {"B", "E"}, {"C", "E"}, {"E", "D"}}

	// visit := make(map[string]bool)
	// adjList["A"] = make([]string, 0)
	// adjList["B"] = make([]string, 0)

	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		if _, ok := adjList[src]; !ok {
			adjList[src] = make([]string, 0)
		}
		if _, ok := adjList[dst]; !ok {
			adjList[dst] = make([]string, 0)
		}
		adjList[src] = append(adjList[src], dst)
	}
	return adjList
}

// Graph Search
// - check if a particular vertex exists in a graph
// - if two vertices are connected
// - to traverse a graph

// Depth First Search
// node: starting node
// target: target node
// visit: keeping track of visited nodes
// counting paths from a node to a target
// Time Complexity: O(V>V)
// Space Complexity: O(V)
func (al *AdjacencyList) DFS(node string, target string, adjList map[string][]string, visit map[string]bool) int {
	if _, ok := visit[node]; ok {
		return 0
	}
	if node == target {
		return 1
	}

	// path from starting node to target node: sequence of edges from starting node to target node
	count := 0

	visit = make(map[string]bool)
	visit[node] = true
	for _, neighbor := range adjList[node] {
		count += al.DFS(neighbor, target, adjList, visit)
	}
	delete(visit, node)
	return count
}

// DFS("A", "E", adjList, set())

// Breadth First Search
// node: starting node
// target: target node
// Shortest path from a node to target
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (al *AdjacencyList) BFS(node string, target string, adjList map[string][]string) int {
	length := 0
	visit := make(map[string]bool)
	q := list.New()
	visit[node] = true
	q.PushBack(node)

	for q.Len() != 0 {
		queueLength := q.Len()
		for i := 0; i < queueLength; i++ {
			curr := q.Front()
			q.Remove(curr)
			if curr.Value == target {
				return length
			}
			for _, neighbor := range adjList[curr.Value.(string)] {
				if _, ok := visit[neighbor]; !ok {
					// add neighbour to visit
					visit[neighbor] = true
					// queue
					q.PushBack(neighbor)
				}
			}
		}
		length++
	}
	return length
}

// BFS("A", "E", adjList)
