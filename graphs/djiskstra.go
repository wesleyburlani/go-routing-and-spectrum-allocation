package graphs

import "sort"

type Djikstra struct{}

func NewDjikstra() *Djikstra {
	return &Djikstra{}
}

func (d Djikstra) FindPaths(
	graph *Graph,
	from *Node,
	to *Node,
	limit int64,
	digraph bool,
) []*Path {
	paths := []*Path{}
	counter := make(map[string]int64)
	for _, node := range graph.Nodes {
		counter[node.Id] = 0
	}

	path := NewPath([]string{from.Id})
	priorityList := []*djikstraNode{newDjikstraNode(from.Id, path, 0)}

	var currentNode *djikstraNode
	for {
		if len(priorityList) == 0 {
			break
		}
		currentNode = priorityList[0]
		priorityList = priorityList[1:]

		if currentNode.Id == to.Id {
			paths = append(paths, currentNode.Path)
			if len(paths) == int(limit) {
				break
			}
			continue
		}

		counter[currentNode.Id]++
		if counter[currentNode.Id] > limit {
			continue
		}

		for _, neighboor := range currentNode.NeighboorsIds(graph, digraph) {
			contains := false
			for _, step := range currentNode.Path.Steps {
				if step == neighboor {
					contains = true
				}
			}

			if contains {
				continue
			}

			newPath := NewPath(currentNode.Path.Steps)
			newPath.Steps = append(newPath.Steps, neighboor)

			hasEdge := false
			var edge *Edge
			for _, l := range graph.Edges {
				if l.To == neighboor {
					hasEdge = true
					edge = l
				}
			}

			if !hasEdge && !digraph {
				for _, l := range graph.Edges {
					if l.From == neighboor {
						edge = l
					}
				}
			}

			cost := currentNode.Distance + edge.Cost
			newNode := newDjikstraNode(neighboor, newPath, cost)
			priorityList = append(priorityList, newNode)
			sortPriorityList(priorityList)
		}
	}
	return paths
}

func sortPriorityList(priorityList []*djikstraNode) {
	sort.SliceStable(priorityList, func(i, j int) bool {
		if priorityList[i].Distance == priorityList[j].Distance {
			return priorityList[i].Id > priorityList[j].Id
		}
		return priorityList[i].Distance > priorityList[j].Distance
	})
}

type djikstraNode struct {
	Id       string
	Path     *Path
	Distance int64
}

func newDjikstraNode(
	Id string,
	path *Path,
	distance int64,
) *djikstraNode {
	return &djikstraNode{
		Id:       Id,
		Path:     path,
		Distance: distance,
	}
}

func (n *djikstraNode) NeighboorsIds(graph *Graph, directional bool) []string {
	var result []string
	for _, edge := range graph.Edges {
		if edge.From == n.Id {
			if directional {
				result = append(result, edge.To)
			} else {
				result = append(result, edge.To)
				result = append(result, edge.From)
			}
		}
	}
	return result
}
