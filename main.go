package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/injections"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/io"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/logs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/options"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/rsa"
)

func loadGraph(nodesFilePath, edgesFilePath string) *graphs.Graph {
	nodesFile := io.OpenFile(nodesFilePath)
	defer nodesFile.Close()

	edgesFile := io.OpenFile(edgesFilePath)
	defer edgesFile.Close()

	return graphs.FromCsv(nodesFile, edgesFile)
}

func main() {

	_, err := flags.Parse(&options.Options)
	if err != nil {
		panic(err)
	}

	graph := loadGraph(options.Options.NodesFilePath, options.Options.EdgesFilePath)

	container := injections.NewContainer(options.Options, graph).BuildContainer()

	container.Invoke(func(a rsa.RSA, l logs.Logger, d demands.Source) {
		demand := d.GetDemands()
		table := a.Start(graph, demand, 8)
		l.Log(table)
	})
}
