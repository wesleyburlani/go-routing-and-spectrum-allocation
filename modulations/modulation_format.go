package modulations

type ModulationFormat interface {
	CapacityInGbps() float64
	ReachabilityInKms() float64
}
