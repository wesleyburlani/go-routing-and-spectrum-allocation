package modulations

type M8QAM struct{}

func (m M8QAM) CapacityInGbps() float64 {
	return 37.5
}

func (m M8QAM) ReachabilityInKms() float64 {
	return 1000
}
