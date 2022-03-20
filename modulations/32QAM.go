package modulations

type M32QAM struct{}

func (m M32QAM) CapacityInGbps() float64 {
	return 62.5
}

func (m M32QAM) ReachabilityInKms() float64 {
	return 250
}
