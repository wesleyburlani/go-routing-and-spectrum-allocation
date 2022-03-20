package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/io"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/rsa"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/utils"
	"go.uber.org/dig"
)

func loadGraph(nodesFilePath, edgesFilePath string) *graphs.Graph {
	nodesFile := io.OpenFile(nodesFilePath)
	defer nodesFile.Close()

	edgesFile := io.OpenFile(edgesFilePath)
	defer edgesFile.Close()

	return graphs.FromCsv(nodesFile, edgesFile)
}

func main() {

	_, err := flags.Parse(&options)
	if err != nil {
		panic(err)
	}

	container := dig.New()
	container.Provide(func() graphs.PathSearch { return graphs.NewDjikstra() })
	container.Provide(func(p graphs.PathSearch) graphs.DisjointedPathPairSearch { return graphs.NewSuurballe(&p) })
	container.Provide(func() rsa.TableFill { return &rsa.FirstFitRsa{} })
	container.Provide(func(t rsa.TableFill) rsa.RSA { return rsa.NewDedicatedProtectionRSA(&t) })

	container.Invoke(func(a rsa.RSA) {
		graph := loadGraph(options.NodesFilePath, options.EdgesFilePath)
		demand := demands.NewGenerator(graph).Generate()
		fmt.Println("demands", len(demand))
		utils.PrintStruct(demand)
		table := a.Start(graph, demand, 8)
		utils.PrintStruct(table.Data)
	})
}
