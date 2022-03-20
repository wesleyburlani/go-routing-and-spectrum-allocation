package graphs

type PathSearch interface {
	FindPaths(
		graph *Graph,
		from *Node,
		to *Node,
		limit int64,
		digraph bool,
	) []*Path
}
