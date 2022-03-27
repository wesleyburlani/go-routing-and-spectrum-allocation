package logs

import (
	"log"
)

type ConsoleLogger struct {
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (s *ConsoleLogger) Log(logs ...interface{}) {
	log.Print(logs...)
}
