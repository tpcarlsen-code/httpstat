package sources

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tpcarlsen-code/mon2http/entities"
)

const tempFile = "/sys/class/thermal/thermal_zone0/temp"

type CPUTemp struct {
	sync.Mutex
	current     entities.Values
	lastUpdated time.Time
	minInterval time.Duration
}

func NewCPUTempSource() *CPUTemp {
	return &CPUTemp{minInterval: time.Second - time.Millisecond}
}

func (c *CPUTemp) Values() entities.Values {
	c.Lock()
	defer c.Unlock()
	if !c.lastUpdated.IsZero() && c.lastUpdated.Add(c.minInterval).After(time.Now()) {
		return c.current
	}
	c.lastUpdated = time.Now()
	raw, _ := os.ReadFile(tempFile)
	t, _ := strconv.Atoi(strings.TrimSpace(string(raw)))
	c.current = entities.Values{
		{
			Name:  "cpu_temp",
			Value: float32(t) / 1000,
		},
	}
	return c.current
}
