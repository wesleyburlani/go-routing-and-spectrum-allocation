package rsa

import (
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
)

type TableFill interface {
	FillDemand(
		table *Table,
		graph *graphs.Graph,
		demand *demands.Demand,
		path *graphs.Path,
		slots []*Slot,
		protection bool,
	) bool
}
