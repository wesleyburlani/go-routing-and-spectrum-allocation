package graphs

import "testing"

func getMockedGraphForDjikstra() *Graph {
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
				From:        "B",
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

func TestDjikstraFindPaths(t *testing.T) {
	graph := getMockedGraphForDjikstra()
	pathSearch := Djikstra{}

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

	paths := pathSearch.FindPaths(graph, nodeA, nodeD, 4, true)

	got := len(paths)
	want := 3

	if int(got) != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
