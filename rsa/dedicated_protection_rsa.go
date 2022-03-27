package rsa

import (
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/logs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/utils"
)

type DedicatedProtectionRSA struct {
	tableFill *TableFill
	logger    *logs.Logger
}

func NewDedicatedProtectionRSA(
	tableFill *TableFill,
	logger *logs.Logger,
) *DedicatedProtectionRSA {
	return &DedicatedProtectionRSA{
		tableFill: tableFill,
		logger:    logger,
	}
}

func (s DedicatedProtectionRSA) Start(
	graph *graphs.Graph,
	demands []*demands.Demand,
	numberOfChannels int,
) *Table {
	var pathSerch graphs.PathSearch = graphs.NewDjikstra()
	disjointedPathSearch := graphs.NewSuurballe(&pathSerch)
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

		paths := disjointedPathSearch.FindDisjointedPaths(graph, from, to, 1, 1)
		if len(paths) == 0 {
			(*s.logger).Log(demand.String())
			continue
		}

		for _, path := range paths {
			tableCopy := &Table{}
			utils.Copy(table, tableCopy)
			(*s.logger).Log("trying main path: ", path.First)
			availableSlots := tableCopy.AvailableSlots(graph, path.First)
			result := (*s.tableFill).FillDemand(tableCopy, graph, demand, path.First, availableSlots, false)
			if result {
				availableSlots2 := tableCopy.AvailableSlots(graph, path.Second)
				(*s.logger).Log("trying secondary path: ", path.Second)
				result = (*s.tableFill).FillDemand(tableCopy, graph, demand, path.Second, availableSlots2, true)
				if result {
					utils.Copy(tableCopy, table)
					(*s.logger).Log(table)
					supplied++
				}
			}
		}
	}
	(*s.logger).Log("total demands: ", len(demands), " supplied: ", supplied, " blocked: ", len(demands)-supplied)
	return table
}
