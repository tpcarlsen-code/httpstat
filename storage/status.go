package storage

import (
	"sync"

	"github.com/tpcarlsen-code/mon2http/entities"
)

type Status struct {
	status entities.Status
	sync.RWMutex
}

func (ss *Status) Set(s entities.Status) {
	ss.Lock()
	defer ss.Unlock()
	ss.status = s
}

func (ss *Status) Get() entities.Status {
	ss.RLock()
	defer ss.RUnlock()
	return ss.status
}
