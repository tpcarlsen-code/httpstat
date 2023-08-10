package storage

import (
	"sync"

	"mon2http/src/entities"
)

type Values struct {
	values entities.Values
	sync.RWMutex
}

func (vs *Values) Set(v entities.Values) {
	vs.Lock()
	defer vs.Unlock()
	vs.values = v
}

func (vs *Values) Get() entities.Values {
	vs.RLock()
	defer vs.RUnlock()
	return vs.values
}
