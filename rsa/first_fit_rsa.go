package rsa

import (
	"strconv"

	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/utils"
)

type FirstFitRsa struct{}

func (f *FirstFitRsa) FillDemand(
	table *Table,
	graph *graphs.Graph,
	demand *demands.Demand,
	path *graphs.Path,
	slots []*Slot,
	protection bool,
) bool {
	edgesPath := graph.GetEdgeArrayFromPath(*path)
	emptySlots := [][]int{}

	var instance *Slot
	for key := range table.Data {
		contains := false
		for _, slot := range slots {
			if slot.Edge.GetId() == key {
				contains = true
				instance = slot
				break
			}
		}
		if contains {
			emptySlots = append(emptySlots, instance.Availables)
		}
	}

	intersection := emptySlots[0]
	for _, slot := range emptySlots[1:] {
		intersection = utils.Intersection(intersection, slot)
	}

	indexesToFill := f.idexesToFill(intersection, int(demand.Slots), 0)

	if indexesToFill == nil {
		return false
	}

	for _, edge := range edgesPath {
		for _, slot := range indexesToFill {
			table.Data[edge.GetId()][slot].ProtectionDemand = protection
			table.Data[edge.GetId()][slot].Values = append(
				table.Data[edge.GetId()][slot].Values,
				strconv.FormatInt(demand.Id, 10),
			)
		}
	}

	return true
}

func (f *FirstFitRsa) idexesToFill(
	intersection []int,
	slotsNumber int,
	iteration int,
) []int {
	if len(intersection) <= iteration ||
		(len(intersection)-iteration) < slotsNumber {
		return nil
	}

	list := []int{intersection[iteration]}
	if len(list) == slotsNumber {
		return list
	}

	for i := iteration + 1; i < len(intersection); i++ {
		if list[len(list)-1] == intersection[i]-1 {
			list = append(list, intersection[i])
		} else {
			break
		}
		if len(list) == slotsNumber {
			return list
		}
	}
	return f.idexesToFill(intersection, slotsNumber, iteration+1)
}
