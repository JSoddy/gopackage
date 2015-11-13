package graph

import (
	"container/heap"
	//"math"
	//"fmt"
)


/* ========================================
A package for graph operations
=========================================*/ 


// Data Definitions
///////////////////////////////////////////////////////

// Node is an element of a Graph
type Node struct {
	Name     	int  // The node name (numbers)
	Distance 	int  // Shortest path Distance to Node
	explored 	bool // True if Node has been explored
	index    	int  // index of the item in the heap
	EdgesOut   	[]Edge
	EdgesIn		[]Edge
}

// Edge is an element of a Graph
type Edge struct {
	Length 	int   // The length of this edge
	Head   	*Node // The node at this Edge's Head
	Tail	*Node // The node at this Edge's Tail
}

// Graph implements heap.Interface and holds Nodes
type Graph struct {
	Vertices	[]*Node
	Edges		[]Edge
}

func (g Graph) Len() int { return len(g.Vertices) }

func (g Graph) Less(i, j int) bool {
	return g.Vertices[i].Distance <= g.Vertices[j].Distance
}

func (g Graph) Swap(i, j int) {
	g.Vertices[i], g.Vertices[j] = g.Vertices[j], g.Vertices[i]
	g.Vertices[i].index = i
	g.Vertices[j].index = j
}

func (g *Graph) Push(x interface{}) {
	n := g.Len()
	node := x.(*Node)
	node.index = n
	g.Vertices = append(g.Vertices, node)
}

func (g *Graph) Pop() interface{} {
	old := *g
	n := old.Len()
	node := old.Vertices[n-1]
	node.index = -1
	g.Vertices = old.Vertices[0 : n-1]
	return node
}

func (g *Graph) update(n *Node, Distance int) {
	if n.Distance > Distance {
		heap.Remove(g, n.index)
		n.Distance = Distance
		heap.Push(g, n)
	}
}

// Function Definitions
///////////////////////////////////////////////////////

// Graph Int -> []int
// computes the shortest path to all elements in Graph g from node n
// using Dijkstra's algorithm
func (g Graph) Dijkstra(n int) []int {
	maxIndex	:= g.Len()
	if maxIndex < 1 {
		return make([]int, 0)
	}
	//set all Node Distance to maxInt
	for i := 0; i < len(g.Vertices); i++ {
		g.Vertices[i].Distance	= 0x3f3f3f3f
		g.Vertices[i].index		= i
		g.Vertices[i].explored 	= false
	}
	//initialize heap
	//set starting Node Distance to 0
	g.update(g.Vertices[n-1], 0)
	heap.Init(&g)
	for g.Len() > 1 {
		node := heap.Pop(&g).(*Node)
		node.explored = true
		for _, edge := range node.EdgesOut {
			if !edge.Head.explored {
				g.update(edge.Head, (edge.Length + node.Distance))
			}
		}
	}
	g.Vertices	= g.Vertices[:maxIndex]
	values		:= make([]int, len(g.Vertices))
	for _, v := range g.Vertices{
		values[v.Name-1]	= v.Distance
	}
	return values
}

// Graph Int -> ()
// computes the shortest path to all elements in Graph g from node n
// using BellmanFord Algorithm
func (g Graph) BellmanFord(n int) ([]int, int) {
	current 	:= make([]int, len(g.Vertices))
	previous	:= make([]int, len(g.Vertices))
	for i, _ := range current {
		current[i] = 0
		//current[i] = 0x3f3f3f3f
		previous[i] = 0x3f3f3f3f
	}
	current[n-1] = 0
	for i := 0; i < len(g.Vertices)+2 && !sameIntSlice(previous, current); i++ {
		current, previous = previous, current
		for j, _ := range current {
			current[j] = bestPath(previous, *g.Vertices[j], j)
		}
	}
	if !sameIntSlice(previous, current) {
		return current, -1
	} else {
		return current, 0
	}
}

func bestPath(previous []int, vertex Node, n int) int {
	currentBest	:= previous[n]
	for _, v := range vertex.EdgesIn {
		thisTry := previous[v.Tail.Name-1] + v.Length
		if thisTry < currentBest {
			currentBest = thisTry
		}
	}
	return currentBest
}

func sameIntSlice(first []int, second []int) bool {
	same 		:= true
	if len(first) != len(second) {
		same =  false
	}
	for i := 0; i < len(first) && same; i++ {
		if first[i] != second[i] {
			same = false
		}
	}
	return same
}