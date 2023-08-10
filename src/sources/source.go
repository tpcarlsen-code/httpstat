package sources

import "mon2http/src/entities"

type Source interface {
	Values() entities.Values
}
