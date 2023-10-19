package sources

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/tpcarlsen-code/mon2http/entities"
)

const cpuStatFile = "/proc/stat"

type AverageCPUSource struct {
	sync.Mutex
	lastTotalUsedCPU int64
	lastTotalIdleCPU int64
	current          entities.Values
	lastUpdated      time.Time
	minInterval      time.Duration
}

func NewAverageCPUSource() *AverageCPUSource {
	return &AverageCPUSource{minInterval: 3*time.Second - time.Millisecond}
}

func (s *AverageCPUSource) Values() entities.Values {
	s.Lock()
	defer s.Unlock()
	now := time.Now()
	if !s.lastUpdated.IsZero() && s.lastUpdated.Add(s.minInterval).After(now) {
		return s.current
	}

	s.lastUpdated = now
	values := getCPUStats()
	if len(values) != 4 {
		log.Printf("WARNING: Wrong len of cpu stats: %d", len(values))
		return nil
	}
	used := values[0] + values[1] + values[2]
	idle := values[3]

	if s.lastTotalUsedCPU == 0 {
		s.lastTotalUsedCPU = used
		s.lastTotalIdleCPU = idle
		return nil
	}
	percentage := float32(used-s.lastTotalUsedCPU) * 100 / float32((used-s.lastTotalUsedCPU)+(idle-s.lastTotalIdleCPU))
	s.lastTotalUsedCPU = used
	s.lastTotalIdleCPU = idle

	s.current = entities.Values{
		{
			Name:  "cpu_total",
			Value: percentage,
		},
	}
	return s.current
}

func getCPUStats() []int64 {
	f, err := os.Open(cpuStatFile)
	defer f.Close()
	if err != nil {
		return nil
	}
	buf := make([]byte, 80) // Enough bytes to read the current stats (first line).
	f.Read(buf)

	return readInts(string(buf), 4)
}
