package graphs

import "testing"

func getMockedGraphForSuurballe() *Graph {
	return &Graph{
		Nodes: []*Node{
			{
				Id:        "A",
				Latitude:  "1",
				Longitude: "1",
				Type:      "A",
			},
			{
				Id:        "B",
				Latitude:  "1",
				Longitude: "1",
				Type:      "B",
			},
			{
				Id:        "C",
				Latitude:  "1",
				Longitude: "1",
				Type:      "C",
			},
			{
				Id:        "D",
				Latitude:  "1",
				Longitude: "1",
				Type:      "D",
			},
		},
		Edges: []*Edge{
			{
				From:        "A",
				To:          "B",
				Length:      1,
				Capacity:    1,
				Cost:        1,
				Designation: "A",
				Delay:       "1",
			},
			{
				From:        "B",
				To:          "C",
				Length:      1,
				Capacity:    1,
				Cost:        1,
				Designation: "A",
				Delay:       "1",
			},
			{
				From:        "C",
				To:          "D",
				Length:      1,
				Capacity:    1,
				Cost:        1,
				Designation: "A",
				Delay:       "1",
			},
			{
				From:        "A",
				To:          "D",
				Length:      1,
				Capacity:    1,
				Cost:        1,
				Designation: "A",
				Delay:       "1",
			},
		},
	}
}

func TestSuurballeFindDisjointedPaths(t *testing.T) {
	graph := getMockedGraphForSuurballe()
	var pathSearch PathSearch = Djikstra{}
	disjointedPathSearch := NewSuurballe(&pathSearch)

	nodeA := &Node{
		Id:        "A",
		Latitude:  "1",
		Longitude: "1",
		Type:      "A",
	}

	nodeD := &Node{
		Id:        "D",
		Latitude:  "1",
		Longitude: "1",
		Type:      "D",
	}

	paths := disjointedPathSearch.FindDisjointedPaths(graph, nodeA, nodeD, 1, 1)

	got := len(paths)
	want := 1

	if int(got) != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if paths[0].First == nil {
		t.Errorf("path First pair Item is nil")
	}

	if paths[0].Second == nil {
		t.Errorf("path Second pair Item is nil")
	}
}
