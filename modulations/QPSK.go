package modulations

type QPSK struct{}

func (m QPSK) CapacityInGbps() float64 {
	return 25
}

func (m QPSK) ReachabilityInKms() float64 {
	return 2000
}
