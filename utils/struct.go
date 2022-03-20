package utils

import (
	"encoding/json"
	"fmt"

	copier "github.com/jinzhu/copier"
)

func PrintStruct(s interface{}) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func Copy(from interface{}, to interface{}) {
	copier.CopyWithOption(to, from, copier.Option{DeepCopy: true})
}
