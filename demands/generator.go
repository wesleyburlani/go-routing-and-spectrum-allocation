package demands

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/logs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/utils"
)

type Generator struct {
	graph        *graphs.Graph
	saveFilePath string
	logger       *logs.Logger
}

func NewGenerator(
	graph *graphs.Graph,
	saveFilePath string,
	logger *logs.Logger,
) *Generator {
	return &Generator{
		graph:        graph,
		saveFilePath: saveFilePath,
		logger:       logger,
	}
}

func (g Generator) GetDemands() []*Demand {
	(*g.logger).Log("generating demands..")
	var mutex sync.Mutex
	var waiter sync.WaitGroup

	demands := []*Demand{}
	ids := g.getNodeIds()
	maxEdges := g.maxNumberOfEdges()
	iterator := utils.RandomNumberBetween(1, maxEdges)
	memory := []string{}

	for i := 0; i < iterator; i++ {
		waiter.Add(1)

		go func(it int) {
			defer waiter.Done()
			from := ids[utils.RandomNumberBetween(0, len(ids)-1)]
			to := ids[utils.RandomNumberBetween(0, len(ids)-1)]

			for {
				if from != to {
					break
				}
				to = ids[utils.RandomNumberBetween(0, len(ids)-1)]
			}

			mutex.Lock()
			hash := from + "-" + to
			contains := false
			for _, el := range memory {
				if el == hash {
					contains = true
					break
				}
			}
			mutex.Unlock()

			if contains {
				return
			}

			demand := utils.RandomNumberBetween(1, 10)
			demandInGbps := 40.0
			mutex.Lock()
			demands = append(demands, NewDemand(int64(it), from, to, demandInGbps, int64(demand)))
			memory = append(memory, hash)
			mutex.Unlock()
		}(i)
	}
	waiter.Wait()
	if g.saveFilePath != "" {
		(*g.logger).Log("saving demands on file..")
		g.saveOnFile(demands)
	}
	(*g.logger).Log(fmt.Sprintf("%d demands generated", len(demands)))
	return demands
}

func (g Generator) getNodeIds() []string {
	ids := []string{}
	for _, node := range g.graph.Nodes {
		ids = append(ids, node.Id)
	}
	return ids
}

func (g Generator) maxNumberOfEdges() int {
	return len(g.graph.Nodes) * (len(g.graph.Nodes) - 1) / 2
}

func (g Generator) saveOnFile(demands []*Demand) {
	f, err := os.Create(g.saveFilePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	stringDemands, _ := utils.Stringfy(demands)
	data := []byte(stringDemands)

	_, err2 := f.Write(data)

	if err2 != nil {
		log.Fatal(err2)
	}
}
