package sources

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"mon2http/src/entities"
)

const uptimeFile = "/proc/uptime"

type Uptime struct {
	sync.Mutex
	current     entities.Values
	lastUpdated time.Time
	minInterval time.Duration
}

func NewUptimeSource() *Uptime {
	return &Uptime{minInterval: time.Minute - time.Millisecond}
}

func (c *Uptime) Values() entities.Values {
	c.Lock()
	defer c.Unlock()
	if !c.lastUpdated.IsZero() && c.lastUpdated.Add(c.minInterval).After(time.Now()) {
		return c.current
	}
	c.lastUpdated = time.Now()
	raw, err := os.ReadFile(uptimeFile)
	if err != nil {
		log.Printf("WARNING: Error reading uptime: %s", err.Error())
		return nil
	}
	days := c.upInDays(string(raw))
	c.current = entities.Values{
		{
			Name:  "uptime",
			Value: days,
		},
	}
	return c.current
}

// 83059.43 322186.03
func (c *Uptime) upInDays(uptimeOutput string) float32 {
	p := strings.Split(uptimeOutput, " ")
	f, _ := strconv.ParseFloat(p[0], 32)
	return float32(f / 86400)
}
