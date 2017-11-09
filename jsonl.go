package jsonl

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
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

// JSONFileRaw reads the .json file and  returns raw JSON as map
func JSONFileRaw(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := bufio.NewReader(file)

	return JSONRaw(buffer)
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

// JSONFileObj reads the .json file and returns raw JSON as map[string]interface{}
func JSONFileObj(path string) (IJSON, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := bufio.NewReader(file)

	return JSONObj(buffer)
}

// Get returns the value by the path (using dot-method) of JSON's hierarchy
// returns `nil` when the `key` has bad format
// returns the `defaultValue` when there is no def. (index out of range) in the json
func (tjson *TJSON) Get(key string, defaultValue interface{}) interface{} {
	var (
		paths                = strings.Split(key, ".")
		re                   = regexp.MustCompile(`(.*)\[(\d+)\]$`)
		thisJSON interface{} = tjson.json
		keys     []string
	)
	for _, path := range paths {
		if len(path) == 0 {
			continue
		}
		if subPath, ok := thisJSON.(map[string]interface{}); ok {
			keys = re.FindStringSubmatch(path)
			if len(keys) != 3 {
				// !array
				if value, exists := subPath[path]; exists {
					thisJSON = value
				} else {
					return defaultValue
				}
			} else {
				if value, exists := subPath[keys[1]]; exists {
					// need arrIdx is type of int
					// as index of array must be int, so when it's not an int,
					// will treat it as an error, nil will return
					arrIdx, err := strconv.Atoi(keys[2])
					if err != nil {
						return nil
					}
					if subPath, ok := value.([]interface{}); ok {
						// need to check the range
						// out of range is not an error
						if len(subPath) <= arrIdx {
							return defaultValue
						}
						thisJSON = subPath[arrIdx]
					} else {
						return nil
					}
				} else {
					return defaultValue
				}
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
