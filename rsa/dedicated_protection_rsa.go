package rsa

import (
	"fmt"

	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/utils"
)

type DedicatedProtectionRSA struct {
	TableFill *TableFill
}

func NewDedicatedProtectionRSA(tableFill *TableFill) *DedicatedProtectionRSA {
	return &DedicatedProtectionRSA{
		TableFill: tableFill,
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

		paths := disjointedPathSearch.FindDisjointedPaths(graph, from, to, 1, 1)
		if len(paths) == 0 {
			continue
		}

		for _, path := range paths {
			tableCopy := &Table{}
			utils.Copy(table, tableCopy)
			availableSlots := tableCopy.AvailableSlots(graph, path.First)
			result := (*s.TableFill).FillDemand(table, graph, demand, path.First, availableSlots, false)
			if result {
				availableSlots = tableCopy.AvailableSlots(graph, path.Second)
				result = (*s.TableFill).FillDemand(table, graph, demand, path.Second, availableSlots, true)
				if result {
					utils.Copy(tableCopy, table)
					fmt.Println("supplied: ", demand.Id)
				}
			}
		}
	}
	return table
}
