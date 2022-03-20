package graphs

import "testing"

func TestGetId(t *testing.T) {

	edge := &Edge{
		From:        "A",
		To:          "B",
		Length:      1,
		Capacity:    1,
		Cost:        1,
		Designation: "A",
		Delay:       "1",
	}

	got := edge.GetId()
	want := "A - B"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
