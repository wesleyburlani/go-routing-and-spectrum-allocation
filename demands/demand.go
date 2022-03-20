package demands

type Demand struct {
	Id           int64   `json:"id"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	DemandInGbps float64 `json:"demand_in_gbps"`
	Slots        int64   `json:"slots"`
}

func NewDemand(
	id int64,
	from string,
	to string,
	demandInGbps float64,
	slots int64,
) *Demand {
	return &Demand{
		Id:           id,
		From:         from,
		To:           to,
		DemandInGbps: demandInGbps,
		Slots:        slots,
	}
}
