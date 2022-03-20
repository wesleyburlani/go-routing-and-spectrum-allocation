package rsa

import (
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
)

type RSA interface {
	Start(graph *graphs.Graph, demands []*demands.Demand, numberOfChannels int) *Table
}
