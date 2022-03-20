package demands

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Reader struct{}

func (r Reader) Read(file *os.File) []*Demand {
	byteValue, _ := ioutil.ReadAll(file)
	d := []*Demand{}
	json.Unmarshal(byteValue, &d)
	return d
}
