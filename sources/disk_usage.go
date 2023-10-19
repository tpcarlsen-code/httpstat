package sources

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/tpcarlsen-code/mon2http/entities"
)

const diskUsageCommand = "lsblk"
const diskUsageOptions = "-rfn"

type diskUsage struct {
	dev            string
	percentageUsed float32
}

type DiskUsage struct {
	sync.Mutex
	current     entities.Values
	lastUpdated time.Time
	minInterval time.Duration
}

func NewDiskUsage() *DiskUsage {
	return &DiskUsage{minInterval: 15*time.Second - time.Millisecond}
}

func (du *DiskUsage) Values() entities.Values {
	du.Lock()
	defer du.Unlock()
	if !du.lastUpdated.IsZero() && du.lastUpdated.Add(du.minInterval).After(time.Now()) {
		return du.current
	}
	du.lastUpdated = time.Now()
	var out entities.Values
	allMounts := du.get()
	if len(allMounts) == 0 {
		log.Println("WARNING: No disk stats")
	}
	for _, values := range allMounts {
		out = append(out, entities.Value{
			Name:   "disk_usage_" + values.dev,
			Metric: `disk_usage{device="` + values.dev + `"}`,
			Value:  values.percentageUsed,
		})
	}
	du.current = out
	return out
}

/*
sda
sda1 ext4 1.0  2e0e3ce1-70b5-4302-92e8-e2610551ae4f 51.9G 2% /storage/disk1
mmcblk0
mmcblk0p1 vfat FAT32 boot 0F92-BECC 205.2M 20% /boot
mmcblk0p2 ext4 1.0 rootfs 41c98998-6a08-4389-bf74-79c9efcf0739 13.8G 48% /
*/

func (du *DiskUsage) get() []diskUsage {
	output, err := exec.Command(diskUsageCommand, diskUsageOptions).Output()
	if err != nil {
		log.Printf("WARNING: Could not read disk stats: %s\n", err.Error())
		return nil
	}
	return du.values(string(output))
}

func (du *DiskUsage) values(stats string) []diskUsage {
	lines := strings.Split(stats, "\n")
	var parts []string
	var percent int64
	var dev string
	var out []diskUsage
	re := regexp.MustCompile("\\s+")
	for _, l := range lines {
		dev = ""
		parts = re.Split(l, -1)
		for _, p := range parts {
			if strings.Index(p, "%") > 0 {
				percent = bInt64([]byte(p))
				dev = parts[0]
			}
		}

		if dev != "" {
			out = append(out, diskUsage{
				dev:            dev,
				percentageUsed: float32(percent),
			})
		}
	}
	return out
}
