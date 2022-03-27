package demands

import (
	"fmt"
)

type Demand struct {
	Id           int64   `json:"id"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	DemandInGbps float64 `json:"demandInGbps"`
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

func (d *Demand) String() string {
	return fmt.Sprintf(
		"Id: %d From: %s To: %s Gbps: %f Slots: %d",
		d.Id, d.From, d.To, d.DemandInGbps, d.Slots,
	)
}
