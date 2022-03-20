package demands

import (
	"sync"

	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/utils"
)

type Generator struct {
	Graph *graphs.Graph
}

func NewGenerator(graph *graphs.Graph) *Generator {
	return &Generator{
		Graph: graph,
	}
}

func (g *Generator) Generate() []*Demand {
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
	return demands
}

func (g *Generator) getNodeIds() []string {
	ids := []string{}
	for _, node := range g.Graph.Nodes {
		ids = append(ids, node.Id)
	}
	return ids
}

func (g *Generator) maxNumberOfEdges() int {
	return len(g.Graph.Nodes) * (len(g.Graph.Nodes) - 1) / 2
}
