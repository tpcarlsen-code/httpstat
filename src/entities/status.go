package entities

import (
	"encoding/json"
	"strings"
)

const (
	StatusOk      = "ok"
	StatusWarning = "warn"
	StatusAlert   = "alert"
)

type Status struct {
	Status string   `json:"status"`
	Alerts []string `json:"alerts"`
}

func (s Status) IsOk() bool {
	return s.Status == StatusOk
}

func (s Status) IsWarning() bool {
	return s.Status == StatusWarning
}

func (s Status) IsAlert() bool {
	return s.Status == StatusAlert
}

func (s Status) Txt() string {
	if s.Status == StatusOk {
		return "STATUS: OK"
	}
	return "ALERT: " + strings.Join(s.Alerts, "\nALERT: ") + "\n"
}

func (s Status) Json() []byte {
	b, _ := json.Marshal(s)
	return b
}
