package graphs

type Path struct {
	Steps []string
}

func NewPath(path []string) *Path {
	p := new(Path)
	p.Steps = path
	return p
}
