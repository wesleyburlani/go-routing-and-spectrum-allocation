package modulations

type BPKS struct{}

func (m BPKS) CapacityInGbps() float64 {
	return 12.5
}

func (m BPKS) ReachabilityInKms() float64 {
	return 4000
}
