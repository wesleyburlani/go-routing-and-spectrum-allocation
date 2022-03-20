package graphs

type Edge struct {
	From        string `csv:"From"`
	To          string `csv:"To"`
	Length      int64  `csv:"Length"`
	Capacity    int64  `csv:"Capacity"`
	Cost        int64  `csv:"Cost"`
	Designation string `csv:"Designation"`
	Delay       string `csv:"Delay"`
}

func (l *Edge) GetId() string {
	return l.From + " - " + l.To
}
