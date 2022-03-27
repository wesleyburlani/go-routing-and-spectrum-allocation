package injections

import (
	"strings"

	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/demands"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/graphs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/io"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/logs"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/options"
	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/rsa"
	"go.uber.org/dig"
)

type Container struct {
	options options.ProgramOptions
	graph   *graphs.Graph
	logger  logs.Logger
}

func NewContainer(
	options options.ProgramOptions,
	graph *graphs.Graph,
) Container {
	logger := buildLogger(options)
	return Container{
		options: options,
		graph:   graph,
		logger:  logger,
	}
}

func (c Container) BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(func() logs.Logger { return c.logger })
	container.Provide(func() demands.Source { return c.buildDemandSource() })
	container.Provide(func() graphs.PathSearch { return c.buildPathSearch() })
	container.Provide(func() graphs.DisjointedPathPairSearch {
		return c.buildDisjointedPathPairSearch()
	})
	container.Provide(func() rsa.TableFill { return c.buildTableFill() })
	container.Provide(func() rsa.RSA { return c.buildRSA() })
	return container
}

func buildLogger(options options.ProgramOptions) logs.Logger {
	if strings.ToLower(options.LogType) == "file" {
		return logs.NewFileLogger(options.LogFilePath)
	}

	if strings.ToLower(options.LogType) == "stdout" {
		return logs.NewConsoleLogger()
	}

	panic("unsupported log type")
}

func (c Container) buildDemandSource() demands.Source {
	if strings.ToLower(c.options.DemandsSource) == "file" {
		c.logger.Log("[CONFIG] using demands from file")
		return demands.NewReader(io.OpenFile(c.options.DemandsFilePath), &c.logger)
	}

	if strings.ToLower(c.options.DemandsSource) == "generate" {
		c.logger.Log("[CONFIG] using demands generator")
		return demands.NewGenerator(c.graph, c.options.DemandsFilePath, &c.logger)
	}

	panic("unsupported demands source")
}

func (c Container) buildPathSearch() graphs.PathSearch {
	if strings.ToLower(c.options.PathSearchAlgorithm) == "djikstra" {
		c.logger.Log("[CONFIG] using djikstra path search")
		return graphs.NewDjikstra()
	}
	panic("path search algorithm not supported")
}

func (c Container) buildDisjointedPathPairSearch() graphs.DisjointedPathPairSearch {
	if strings.ToLower(c.options.DisjointedPathPairSearchAlgorithm) == "suurballe" {
		c.logger.Log("[CONFIG] using suurballe disjointed path pair search")
		pathSearch := c.buildPathSearch()
		return graphs.NewSuurballe(&pathSearch)
	}
	panic("unsupported disjointed path pair search algorithm")
}

func (c Container) buildTableFill() rsa.TableFill {
	if strings.ToLower(c.options.TableFillAlgorithm) == "first_fit_rsa" {
		c.logger.Log("[CONFIG] using first fit rsa table fill")
		return &rsa.FirstFitRsa{}
	}

	if strings.ToLower(c.options.TableFillAlgorithm) == "first_fit_rmlsa" {
		c.logger.Log("[CONFIG] using first fit rmlsa table fill")
		return &rsa.FirstFitRmlsa{}
	}

	panic("unsupported table fill algorithm")
}

func (c Container) buildRSA() rsa.RSA {
	tableFill := c.buildTableFill()
	logger := c.logger
	if strings.ToLower(c.options.RsaType) == "single" {
		c.logger.Log("[CONFIG] using single rsa")
		return rsa.NewSingleRSA(&tableFill, &logger)
	}
	if strings.ToLower(c.options.RsaType) == "dedicated_protection" {
		c.logger.Log("[CONFIG] using dedicated protection rsa")
		return rsa.NewDedicatedProtectionRSA(&tableFill, &logger)
	}
	if strings.ToLower(c.options.RsaType) == "shared_protection" {
		c.logger.Log("[CONFIG] using shared protection rsa")
		panic("shared_protection is not supported")
	}
	panic("unsupported rsa type")
}
