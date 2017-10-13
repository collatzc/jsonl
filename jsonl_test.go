package jsonl

import (
	"bytes"
	"testing"
)

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
	t.Error(j)
	t.Error(j["root"])
}

func TestJsonObj(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"root" : {
				"key" : "abc"
			}
		}
	`))
	j, err := JSONObj(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Error(j)
	t.Error(j.Get("root.key", "123"))
}
