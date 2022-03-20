package graphs

import "testing"

func getMockedGraph() *Graph {
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
		},
	}
}

func TestChangeEdgeCost(t *testing.T) {
	graph := getMockedGraph()
	graph.ChangeEdgeCost("A", "B", 2)
	got := graph.Edges[0].Cost
	want := 2

	if int(got) != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRemoveEdge(t *testing.T) {
	graph := getMockedGraph()
	graph.RemoveEdge("A", "B")
	got := len(graph.Edges)
	want := 0

	if int(got) != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetEdgeArrayFromPath(t *testing.T) {
	graph := getMockedGraph()
	path := Path{
		Steps: []string{"A", "B"},
	}
	edgeArray := graph.GetEdgeArrayFromPath(path)
	got := len(edgeArray)
	want := 1

	if int(got) != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
