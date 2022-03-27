package demands

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/wesleyburlani/go-routing-and-spectrum-allocation/logs"
)

type Reader struct {
	file   *os.File
	logger *logs.Logger
}

func NewReader(file *os.File, logger *logs.Logger) *Reader {
	return &Reader{
		file,
		logger,
	}
}

func (r Reader) GetDemands() []*Demand {
	(*r.logger).Log("loading demands from file")
	byteValue, _ := ioutil.ReadAll(r.file)
	d := []*Demand{}
	json.Unmarshal(byteValue, &d)
	return d
}
