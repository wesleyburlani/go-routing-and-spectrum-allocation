package demands

import "github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"

type PathPair struct {
	First  *graphs.Path
	Second *graphs.Path
}

type EdgePair struct {
	Demand   *Demand
	PathPair *PathPair
}

func NewEdgePair(demand *Demand, pathPair *PathPair) *EdgePair {
	return &EdgePair{
		Demand:   demand,
		PathPair: pathPair,
	}
}
