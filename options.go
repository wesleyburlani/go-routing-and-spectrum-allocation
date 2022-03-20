package main

var options struct {
	NodesFilePath string `long:"nodesFilePath" description:"file path of the nodes specifications" required:"true"`
	EdgesFilePath string `long:"edgesFilePath" description:"file path of the edges specifications" required:"true"`
}
