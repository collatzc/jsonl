package jsonl

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

// IJSON is the interface that wraps the basic JSON object
type IJSON interface {
	Get(key string, defaultValue interface{}) interface{}
}

// TJSON is the data type to store the JSON
type TJSON struct {
	json map[string]interface{}
}

// JSONRaw returns the raw JSON as map
func JSONRaw(reader io.Reader) (map[string]interface{}, error) {
	jsonBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}

// JSONObj returns the objective JSON
func JSONObj(reader io.Reader) (IJSON, error) {
	jsonBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
		return nil, err
	}

	return &TJSON{jsonData}, nil
}

// Get returns the value by the path of JSON's hierarchy
func (tjson *TJSON) Get(key string, defaultValue interface{}) interface{} {
	paths := strings.Split(key, ".")
	var thisJSON interface{} = tjson.json
	for _, path := range paths {
		if len(path) == 0 {
			continue
		}
		if subPath, ok := thisJSON.(map[string]interface{}); ok {
			if value, exists := subPath[path]; exists {
				thisJSON = value
			} else {
				return defaultValue
			}
		} else if subPath, ok := thisJSON.(map[interface{}]interface{}); ok {
			if value, exists := subPath[path]; exists {
				thisJSON = value
			} else {
				return defaultValue
			}
		} else {
			return defaultValue
		}
	}
	return thisJSON
}
