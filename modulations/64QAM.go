package modulations

type M64QAM struct{}

func (m M64QAM) CapacityInGbps() float64 {
	return 75
}

func (m M64QAM) ReachabilityInKms() float64 {
	return 125
}
