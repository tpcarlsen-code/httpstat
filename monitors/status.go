package monitors

import (
	"github.com/tpcarlsen-code/mon2http/entities"
)

type Status struct {
	Ok      bool
	Value   entities.Value
	Type    string
	Message string
}

type Alert string
