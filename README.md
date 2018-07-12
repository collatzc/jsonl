# jsonl(eser) ![version 0.1.0](https://img.shields.io/badge/version-0.1.0-brightgreen.svg)

A simple JSON-Object generator

## Features

* Support dot-method, just like access the member variables
* Support using index to get the element of array

## Installation

```
go get github.com/collatzc/jsonl
```

## Usage

Using `JSONObj` to generate a JSON-Object:

```golang
reader := bytes.NewReader([]byte(`
    {
        "root" : {
            "key" : "abc"
        }
    }
`))

j, err := JSONObj(reader)

# get the value of "key"
key := j.Get("root.key")
```

Work well with array:

```golang
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

# get the value of the first "key"
key := j.Get("root[1].key")

```
