package monitors

import "github.com/tpcarlsen-code/mon2http/sources"

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
