package graphs

type Node struct {
	Id        string `csv:"Id"`
	Latitude  string `csv:"Lat"`
	Longitude string `csv:"Long"`
	Type      string `csv:"Type"`
}
