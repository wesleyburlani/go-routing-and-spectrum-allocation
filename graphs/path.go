package graphs

import "strings"

type Path struct {
	Steps []string
}

func NewPath(path []string) *Path {
	p := new(Path)
	p.Steps = path
	return p
}

func (p *Path) String() string {
	return strings.Join(p.Steps, " -> ")
}
