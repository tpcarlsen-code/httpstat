package monitors

import (
	"log"

	"mon2http/src/entities"
)

var availableMonitors []monitor

func registerMonitor(m monitor) {
	availableMonitors = append(availableMonitors, m)
}

type monitor interface {
	Init(numSamples int)
	Status() []Status
	Type() string
}

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Init(numSamples int) {
	for _, mon := range availableMonitors {
		mon.Init(numSamples)
	}
}

func (m *Manager) Update() ([]Alert, []entities.Value) {
	var allValues []entities.Value
	var alerts []Alert
	for _, mon := range availableMonitors {
		statuses := mon.Status()
		for _, s := range statuses {
			if !s.Ok {
				log.Println("ALERT: " + s.Message)
				alerts = append(alerts, Alert(s.Message))
			}
			allValues = append(allValues, s.Value)
		}
	}
	return alerts, allValues
}
