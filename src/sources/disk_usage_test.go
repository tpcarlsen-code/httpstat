package sources

import (
	"testing"
)

func TestDiskUsage_Values(t *testing.T) {
	du := NewDiskUsage()

	raw := `sda       
sda1 ext4 1.0  2e0e3ce1-70b5-4302-92e8-e2610551ae4f 51.9G 2% /storage/disk1
mmcblk0       
mmcblk0p1 vfat FAT32 boot 0F92-BECC 205.2M 20% /boot
mmcblk0p2 ext4 1.0 rootfs 41c98998-6a08-4389-bf74-79c9efcf0739 13.8G 48% /`

	res := du.values(raw)
	if len(res) != 3 {
		t.Errorf("expected len 3, got %d", len(res))
	}
}
