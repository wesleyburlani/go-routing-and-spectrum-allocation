package demands

type Source interface {
	GetDemands() []*Demand
}
