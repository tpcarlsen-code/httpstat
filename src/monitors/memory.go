package monitors

import "mon2http/src/sources"

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
