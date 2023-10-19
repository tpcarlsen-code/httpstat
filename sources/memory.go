package sources

import (
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/tpcarlsen-code/mon2http/entities"
)

const memoryCMD = "free"

type Memory struct {
	sync.Mutex
	current     entities.Values
	lastUpdated time.Time
	minInterval time.Duration
}

func NewMemorySource() *Memory {
	return &Memory{minInterval: time.Second - time.Millisecond}
}

func (c *Memory) Values() entities.Values {
	c.Lock()
	defer c.Unlock()
	if !c.lastUpdated.IsZero() && c.lastUpdated.Add(c.minInterval).After(time.Now()) {
		return c.current
	}
	c.lastUpdated = time.Now()
	output, err := exec.Command(memoryCMD).Output()
	if err != nil {
		log.Printf("WARNING: Error reading mem: %s", err.Error())
		return nil
	}
	lines := strings.Split(string(output), "\n")
	lines = lines[1:]
	memStats := readInts(lines[0], 6)
	if len(memStats) != 6 {
		log.Printf("WARNING: Wrong len of mem stats: %d", len(memStats))
		return nil
	}
	c.current = entities.Values{
		{
			Name:  "memory_usage",
			Value: float32(memStats[5]) * 100.0 / float32(memStats[0]),
		},
	}
	return c.current
}
