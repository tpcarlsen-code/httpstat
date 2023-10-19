package entities

import (
	"encoding/json"
	"fmt"
)

type Value struct {
	Name   string  `json:"name"`
	Metric string  `json:"metric,omitempty"`
	Value  float32 `json:"value"`
}

type Values []Value

func (v Values) Json() []byte {
	b, _ := json.Marshal(v)
	return b
}

func (v Values) Txt() (s string) {
	for _, value := range v {
		s += fmt.Sprintf("%s: %v\n", value.Name, value.Value)
	}
	return
}

func (v Values) Metrics() (s string) {
	var n string
	for _, value := range v {
		n = value.Name
		if value.Metric != "" {
			n = value.Metric
		}
		s += fmt.Sprintf("%s %v\n", n, value.Value)
	}
	return
}
