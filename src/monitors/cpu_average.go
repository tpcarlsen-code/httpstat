package monitors

import "mon2http/src/sources"

const averageCPUType = "cpu_average"
const cpuAverageLimit = 80 // milli-percent

func init() {
	registerMonitor(NewAverageCPU(sources.NewAverageCPUSource()))
}

type AverageCPU struct {
	BaseMonitor
}

func NewAverageCPU(s *sources.AverageCPUSource) *AverageCPU {
	a := AverageCPU{}
	a.source = s
	a.limit = cpuAverageLimit
	a.typeString = averageCPUType
	return &a
}
