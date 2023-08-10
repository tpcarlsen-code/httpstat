package monitors

import (
	"mon2http/src/entities"
)

type Status struct {
	Ok      bool
	Value   entities.Value
	Type    string
	Message string
}

type Alert string
