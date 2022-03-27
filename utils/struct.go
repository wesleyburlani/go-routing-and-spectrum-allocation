package utils

import (
	"encoding/json"

	copier "github.com/jinzhu/copier"
)

func Stringfy(s interface{}) (string, error) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Copy(from interface{}, to interface{}) {
	copier.CopyWithOption(to, from, copier.Option{DeepCopy: true})
}
