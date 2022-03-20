package graphs

type PathPair struct {
	First  *Path
	Second *Path
}

type DisjointedPathPairSearch interface {
	FindDisjointedPaths(
		graph *Graph,
		from *Node,
		to *Node,
		mainPathsLimit int64,
		secondayPathsLimit int64,
	) []*PathPair
}
