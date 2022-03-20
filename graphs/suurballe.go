package graphs

import (
	utils "github.com/wesleyburlani/go-routing-and-spectrum-allocation/utils"
)

type Suurballe struct {
	pathSearch *PathSearch
}

func NewSuurballe(pathSearch *PathSearch) *Suurballe {
	return &Suurballe{pathSearch: pathSearch}
}

func (s Suurballe) FindDisjointedPaths(
	graph *Graph,
	from *Node,
	to *Node,
	mainPathsLimit int64,
	secondayPathsLimit int64,
) []*PathPair {
	digraph := &Graph{}
	utils.Copy(graph, digraph)
	buildDigraphEdges(digraph)

	disjointedPaths := []*PathPair{}
	mainPaths := (*s.pathSearch).FindPaths(graph, from, to, mainPathsLimit, true)

	for _, mainPath := range mainPaths {
		digraphCopy := &Graph{}
		utils.Copy(digraph, digraphCopy)
		for i := 0; i < len(mainPath.Steps)-1; i++ {
			step := mainPath.Steps[i]
			nextStep := mainPath.Steps[i+1]
			digraphCopy.RemoveEdge(step, nextStep)
			digraphCopy.ChangeEdgeCost(nextStep, step, 0)
		}

		secondaryPaths := (*s.pathSearch).FindPaths(digraphCopy, from, to, secondayPathsLimit, true)

		if len(secondaryPaths) == 0 {
			panic("It was not possible to find two disjointed paths between " + from.Id + " and " + to.Id)
		}

		for _, secondaryPath := range secondaryPaths {
			ringSubGraph := createRingSubGraph(digraph, mainPath)
			changeEdgesExistenceState(digraph, secondaryPath, ringSubGraph)
			path := (*s.pathSearch).FindPaths(ringSubGraph, from, to, 1, true)
			if len(path) == 0 {
				continue
			}
			for i := 0; i < len(path[0].Steps)-1; i++ {
				step := path[0].Steps[i]
				nextStep := path[0].Steps[i+1]
				ringSubGraph.RemoveEdge(step, nextStep)
				ringSubGraph.RemoveEdge(nextStep, step)
			}
			path2 := (*s.pathSearch).FindPaths(ringSubGraph, from, to, 1, true)
			if len(path2) > 0 {
				disjointedPaths = append(disjointedPaths, &PathPair{First: path[0], Second: path2[0]})
			}
		}
	}
	return disjointedPaths
}

func buildDigraphEdges(graph *Graph) {
	for _, edge := range graph.Edges {
		newEdge := &Edge{}
		utils.Copy(edge, newEdge)
		newEdge.From = edge.To
		newEdge.To = edge.From
		graph.Edges = append(graph.Edges, newEdge)
	}
}

func createRingSubGraph(graph *Graph, path *Path) *Graph {
	ringSubGraph := &Graph{}
	utils.Copy(graph, ringSubGraph)

	ringSubGraph.Edges = []*Edge{}
	for i := 0; i < len(path.Steps)-1; i++ {
		step := path.Steps[i]
		nextStep := path.Steps[i+1]
		for _, edge := range graph.Edges {
			if edge.From == step && edge.To == nextStep {
				edgeCopy := &Edge{}
				utils.Copy(edge, edgeCopy)
				ringSubGraph.Edges = append(ringSubGraph.Edges, edgeCopy)
				break
			}
		}
	}

	buildDigraphEdges(ringSubGraph)
	return ringSubGraph
}

func changeEdgesExistenceState(
	graph *Graph,
	secondaryPath *Path,
	ringSubGraph *Graph,
) {
	for i := 0; i < len(secondaryPath.Steps)-1; i++ {
		step := secondaryPath.Steps[i]
		nextStep := secondaryPath.Steps[i+1]
		hasEdge := false
		for _, edge := range ringSubGraph.Edges {
			if edge.From == step && edge.To == nextStep {
				hasEdge = true
				break
			}
		}

		if hasEdge {
			ringSubGraph.RemoveEdge(step, nextStep)
			ringSubGraph.RemoveEdge(nextStep, step)
		} else {
			edge := &Edge{}
			for _, e := range graph.Edges {
				if e.From == step && e.To == nextStep {
					edge = e
					break
				}
			}

			edgeCopy := &Edge{}
			utils.Copy(&edge, &edgeCopy)
			edgeCopy2 := &Edge{}
			utils.Copy(&edge, &edgeCopy2)
			edgeCopy2.From = edge.To
			edgeCopy2.To = edge.From
			ringSubGraph.Edges = append(ringSubGraph.Edges, edgeCopy2)
			ringSubGraph.Edges = append(ringSubGraph.Edges, edgeCopy)
		}
	}
}
