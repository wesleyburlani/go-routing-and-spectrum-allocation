package rsa

import (
	"math"
	"sort"
	"strconv"

	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/modulations"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/utils"
)

type FirstFitRmlsa struct{}

func (f *FirstFitRmlsa) FillDemand(
	table *Table,
	graph *graphs.Graph,
	demand *demands.Demand,
	path *graphs.Path,
	slots []*Slot,
	protection bool,
) bool {
	edgesPath := graph.GetEdgeArrayFromPath(*path)

	var totalDistance int64 = 0
	for _, edge := range edgesPath {
		totalDistance += edge.Length
	}

	var format *modulations.ModulationFormat
	multiplier := 1
	for {
		format = f.modulationFormat(demand, float64(totalDistance), multiplier)
		multiplier = multiplier + 1
		if format != nil {
			break
		}
	}

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

	slotsNumber := int(math.Ceil(demand.DemandInGbps / float64(table.EdgeCapacity)))
	indexesToFill := f.idexesToFill(intersection, slotsNumber, 0)

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
			table.Data[edge.GetId()][slot].ModulationFormat = format
		}
	}

	return true
}

func (f *FirstFitRmlsa) modulationFormat(
	demand *demands.Demand,
	distance float64,
	multiplier int,
) *modulations.ModulationFormat {
	formats := []modulations.ModulationFormat{
		&modulations.BPKS{},
		&modulations.QPSK{},
		&modulations.M8QAM{},
		&modulations.M16QAM{},
		&modulations.M32QAM{},
		&modulations.M64QAM{},
	}
	candidates := []modulations.ModulationFormat{}
	for _, format := range formats {
		if format.ReachabilityInKms()*float64(multiplier) >= distance {
			candidates = append(candidates, format)
		}
	}

	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i].ReachabilityInKms() >= candidates[j].ReachabilityInKms()
	})

	return &candidates[0]
}

func (f *FirstFitRmlsa) idexesToFill(
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
