package rsa

import (
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/modulations"
)

type TableData struct {
	ProtectionDemand bool
	Values           []string
	ModulationFormat *modulations.ModulationFormat
}

func NewTableData(
	protectionDemand bool,
	values []string,
) *TableData {
	return &TableData{
		ProtectionDemand: protectionDemand,
		Values:           values,
	}
}

type Table struct {
	Data             map[string]map[int]*TableData
	Graph            *graphs.Graph
	NumberOfChannels int
	EdgeCapacity     float64
}

func NewTable(
	graph *graphs.Graph,
	numberOfChannels int,
	edgeCapacity float64,
) *Table {
	table := &Table{
		Data:             make(map[string]map[int]*TableData),
		Graph:            graph,
		NumberOfChannels: numberOfChannels,
		EdgeCapacity:     edgeCapacity,
	}

	for _, edge := range graph.Edges {
		table.Data[edge.GetId()] = make(map[int]*TableData)
		for i := 0; i < numberOfChannels; i++ {
			table.Data[edge.GetId()][i] = NewTableData(
				false,
				[]string{},
			)
		}
	}
	return table
}

func (t *Table) AvailableSlots(graph *graphs.Graph, path *graphs.Path) []*Slot {
	var availableSlots []*Slot
	edgesPath := graph.GetEdgeArrayFromPath(*path)
	for _, edge := range edgesPath {
		available := []int{}
		for key, element := range t.Data[edge.GetId()] {
			if len(element.Values) == 0 {
				available = append(available, key)
			}
		}
		availableSlots = append(availableSlots, NewSlot(edge, available))
	}
	return availableSlots
}
