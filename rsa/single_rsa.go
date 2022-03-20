package rsa

import (
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
)

type SingleRSA struct {
	TableFill *TableFill
}

func NewSingleRSA(tableFill *TableFill) *SingleRSA {
	return &SingleRSA{
		TableFill: tableFill,
	}
}

func (s SingleRSA) Start(
	graph *graphs.Graph,
	demands []*demands.Demand,
	numberOfChannels int,
) *Table {
	pathSercher := graphs.NewDjikstra()
	table := NewTable(graph, numberOfChannels, 12.5)

	for _, demand := range demands {
		var from *graphs.Node
		var to *graphs.Node
		for _, node := range graph.Nodes {
			if node.Id == demand.From {
				from = node
			}
			if node.Id == demand.To {
				to = node
			}
		}

		paths := pathSercher.FindPaths(graph, from, to, 2, false)
		if len(paths) == 0 {
			continue
		}

		for _, path := range paths {
			availableSlots := table.AvailableSlots(graph, path)
			result := (*s.TableFill).FillDemand(table, graph, demand, path, availableSlots, false)
			if result {
				break
			}
		}
	}
	return table
}
