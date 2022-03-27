package rsa

import (
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/logs"
)

type SingleRSA struct {
	tableFill *TableFill
	logger    *logs.Logger
}

func NewSingleRSA(
	tableFill *TableFill,
	logger *logs.Logger,
) *SingleRSA {
	return &SingleRSA{
		tableFill,
		logger,
	}
}

func (s SingleRSA) Start(
	graph *graphs.Graph,
	demands []*demands.Demand,
	numberOfChannels int,
) *Table {
	pathSercher := graphs.NewDjikstra()
	table := NewTable(graph, numberOfChannels, 12.5)

	supplied := 0
	for _, demand := range demands {
		(*s.logger).Log("processing demand", demand)
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
			(*s.logger).Log("trying path: ", path)
			availableSlots := table.AvailableSlots(graph, path)
			result := (*s.tableFill).FillDemand(table, graph, demand, path, availableSlots, false)
			if result {
				supplied++
				(*s.logger).Log("demand supplied")
				break
			}
			(*s.logger).Log("demand not supplied")
		}
	}
	(*s.logger).Log("total demands: ", len(demands), " supplied: ", supplied, " blocked: ", len(demands)-supplied)
	return table
}
