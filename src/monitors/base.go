package monitors

import (
	"fmt"

	"mon2http/src/sources"
)

type BaseMonitor struct {
	source     sources.Source
	numSamples int
	limit      float32
	typeString string

	values map[string][]float32
}

func (m *BaseMonitor) Init(numSamples int) {
	m.numSamples = numSamples
}

func (m *BaseMonitor) Status() []Status {
	if m.values == nil {
		m.values = map[string][]float32{}
	}
	values := m.source.Values()
	if len(values) == 0 {
		return nil
	}
	var out []Status
	var s Status
	var current float32
	for _, v := range values {
		s = Status{Ok: true, Type: m.Type()}
		m.values[v.Name] = append(m.values[v.Name], v.Value)

		if len(m.values[v.Name]) >= m.numSamples {
			m.values[v.Name] = trimLeft(m.values[v.Name], m.numSamples)
			current = avg(m.values[v.Name])

			if m.limit > 0 && current >= m.limit {
				s.Message = fmt.Sprintf(
					"%s: %s: %f higher than limit %f",
					m.Type(), v.Name, current, m.limit,
				)
				s.Ok = false
			}
		}

		s.Value = v
		out = append(out, s)
	}
	return out
}

func (m *BaseMonitor) Type() string {
	return m.typeString
}
