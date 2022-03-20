package rsa

import "github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"

type Slot struct {
	Edge       *graphs.Edge
	Availables []int
}

func NewSlot(edge *graphs.Edge, availables []int) *Slot {
	return &Slot{
		Edge:       edge,
		Availables: availables,
	}
}
