package graphs

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

// NewGraph creates a new Graph instance
func NewGraph(Nodes []*Node, Edges []*Edge) *Graph {
	g := new(Graph)
	g.Nodes = Nodes
	g.Edges = Edges
	return g
}

// FromCsv parses 2 CSV input files into a Graph
func FromCsv(nodesFile *os.File, edgesFile *os.File) *Graph {
	nodes := []*Node{}
	if err := gocsv.UnmarshalFile(nodesFile, &nodes); err != nil {
		fmt.Println("Error unmarshalling nodes:", err)
		os.Exit(1)
	}

	edges := []*Edge{}
	if err := gocsv.UnmarshalFile(edgesFile, &edges); err != nil {
		fmt.Println("Error unmarshalling edges:", err)
		os.Exit(1)
	}

	return NewGraph(nodes, edges)
}

// ChangeEdgeCost changes the cost of a edge
func (graph *Graph) ChangeEdgeCost(nodeFrom string, nodeTo string, cost int64) {
	for _, edge := range graph.Edges {
		if edge.From == nodeFrom && edge.To == nodeTo {
			edge.Cost = cost
		}
	}
}

// RemoveEdge removes a edge from the Graph
func (graph *Graph) RemoveEdge(nodeFrom string, nodeTo string) {
	for i, edge := range graph.Edges {
		if edge.From == nodeFrom && edge.To == nodeTo {
			graph.Edges = append(graph.Edges[:i], graph.Edges[i+1:]...)
		}
	}
}

// GetEdgeArrayFromPath returns an array of edges based on a path
func (graph *Graph) GetEdgeArrayFromPath(path Path) []*Edge {
	var result []*Edge
	for i := 0; i < len(path.Steps)-1; i++ {
		for _, edge := range graph.Edges {
			if (edge.From == path.Steps[i] && edge.To == path.Steps[i+1]) ||
				(edge.To == path.Steps[i] && edge.From == path.Steps[i+1]) {
				result = append(result, edge)
			}
		}
	}
	return result
}
