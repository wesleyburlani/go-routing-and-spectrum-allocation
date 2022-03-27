package rsa

import (
	"strings"

	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/utils"
)

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

func (s *Slot) String() string {
	return "slot(" + s.Edge.GetId() + "): " + strings.Join(utils.IntSliceToStringSlice(s.Availables), ", ")
}
