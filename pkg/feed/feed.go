package feed

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Storage map[string]Feed

type Feed struct {
	Identifier string   `json:"identifier"`
	Title      string   `json:"title"`
	Links      []string `json:"links"`
}

func Save(path string, feeds Storage) error {
	data, err := json.Marshal(feeds)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, os.ModePerm)
}

func Load(path string) (Storage, error) {
	var feeds Storage

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &feeds)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}
