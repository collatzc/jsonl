package jsonl

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestJSONFileRaw(t *testing.T) {
	j, err := JSONFileRaw("./test.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(j["root"])
}

func TestJSONFileObj(t *testing.T) {
	j, err := JSONFileObj("./test.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(j.Get("root.key", "default"))
	if j.Get("root.key", "default") != "abc" {
		t.Error("someth. wrong with JSONFileObj()")
	}
}

func TestJSONRaw(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"root" : {
				"key" : "abc"
			}
		}
	`))
	j, err := JSONRaw(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if tmp := j["root"].(map[string]interface{}); tmp["key"] != "abc" {
		t.Error(j["root"])
		t.Fail()
	}
}

func TestJsonObj(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"root": [
				{
					"key": "abc"
				},
				{
					"key": [
						{
							"1": "def"
						}
					]
				}
			]
		}
	`))
	j, err := JSONObj(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(j.Get("root[1].key[0].1", "123"))
	if j.Get("root[1].key[0].1", "123") != "def" {
		t.Error("dot-method with array has some problem")
	}

	file, err := os.Open("./test.json")
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	j2, err := JSONObj(buf)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if j2.Get("root.key", "123") != "abc" {
		t.Error("something wrong with parse from file with the dot-method")
	}
}

func TestRegex(t *testing.T) {
	re := regexp.MustCompile(`(.*)\[(\d+)\]`)
	keys := re.FindStringSubmatch("abc")
	t.Log(keys)
}
