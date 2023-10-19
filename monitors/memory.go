package monitors

import "github.com/tpcarlsen-code/mon2http/sources"

const memoryUsageType = "memory_usage"
const memoryUsageLimit = 80 //percent

func init() {
	registerMonitor(NewMemoryUsage(sources.NewMemorySource()))
}

type MemoryUsage struct {
	BaseMonitor
}

func NewMemoryUsage(s *sources.Memory) *MemoryUsage {
	m := MemoryUsage{}
	m.source = s
	m.limit = memoryUsageLimit
	m.typeString = memoryUsageType
	return &m
}
