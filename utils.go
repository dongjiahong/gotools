package gotools

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func DecodeJsonFromFile(jsonFile string, schema interface{}) error {
	f, err := os.Open(jsonFile)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	return dec.Decode(schema)
}
