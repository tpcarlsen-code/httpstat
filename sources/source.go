package sources

import "github.com/tpcarlsen-code/mon2http/entities"

type Source interface {
	Values() entities.Values
}
