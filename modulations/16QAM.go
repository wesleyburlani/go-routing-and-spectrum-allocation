package modulations

type M16QAM struct{}

func (m M16QAM) CapacityInGbps() float64 {
	return 50
}

func (m M16QAM) ReachabilityInKms() float64 {
	return 500
}
