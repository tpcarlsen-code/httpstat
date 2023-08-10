package sources

import (
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"mon2http/src/entities"
)

const diskStatFile = "/proc/diskstats"

type diskStat struct {
	timestamp             time.Time
	dev                   string
	readBytes, writeBytes int64
}

type DiskRW struct {
	sync.Mutex
	lastStats   map[string]diskStat
	current     entities.Values
	lastUpdated time.Time
	minInterval time.Duration
}

func NewDiskRW() *DiskRW {
	return &DiskRW{
		lastStats:   map[string]diskStat{},
		minInterval: 3*time.Second - time.Millisecond,
	}
}

func (d *DiskRW) Values() entities.Values {
	d.Lock()
	defer d.Unlock()
	now := time.Now()
	if !d.lastUpdated.IsZero() && d.lastUpdated.Add(d.minInterval).After(now) {
		return d.current
	}

	d.lastUpdated = now
	stats := d.getDiskStats()
	if len(d.lastStats) == 0 {
		d.lastStats = stats
		return nil
	}

	var values entities.Values
	var rbs, wbs float32
	for dev, stat := range stats {
		lastStat, ok := d.lastStats[dev]
		if !ok {
			log.Printf("WARN: No previous stats found for device %s\n", dev)
			continue
		}
		dur := stat.timestamp.Sub(lastStat.timestamp)
		rbs = float32(((stat.readBytes - lastStat.readBytes) * 1000) / dur.Milliseconds())
		wbs = float32(((stat.writeBytes - lastStat.writeBytes) * 1000) / dur.Milliseconds())
		values = append(values, entities.Value{
			Name:   "disk_read_" + dev,
			Metric: `disk_read{device="` + dev + `"}`,
			Value:  rbs,
		})
		values = append(values, entities.Value{
			Name:   "disk_write_" + dev,
			Metric: `disk_write{device="` + dev + `"}`,
			Value:  wbs,
		})
	}
	d.lastStats = stats
	d.current = values
	return d.current
}

func (d *DiskRW) getDiskStats() map[string]diskStat {
	f, err := os.ReadFile(diskStatFile)
	if err != nil {
		log.Println("ERROR: Could not read /proc/diskstats")
		return nil
	}
	return d.parseDiskStats(string(f))
}

func (d *DiskRW) parseDiskStats(statsData string) map[string]diskStat {
	rTime := time.Now()
	stats := map[string]diskStat{}
	lines := strings.Split(statsData, "\n")
	re := regexp.MustCompile("\\s+")
	var parts []string
	var dev string
	var bRead, bWrite int64
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts = re.Split(strings.TrimSpace(line), -1)
		if parts[0] != "1" && parts[0] != "7" && parts[1] == "0" { // major device
			dev = parts[2]
			bRead = bInt64([]byte(parts[5])) * 512
			bWrite = bInt64([]byte(parts[9])) * 512
			stats[dev] = diskStat{
				timestamp:  rTime,
				dev:        dev,
				readBytes:  bRead,
				writeBytes: bWrite,
			}
		}
	}
	return stats
}
