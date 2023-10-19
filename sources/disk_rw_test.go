package sources

import (
	"testing"
)

const test = `179       0 mmcblk0 18440 5008 1447882 75470 95669 64547 1458338 2964515 0 295432 3039985 0 0 0 0 0 0
179       1 mmcblk0p1 230 57 11168 350 3 0 10 8 0 384 358 0 0 0 0 0 0
179       2 mmcblk0p2 18174 4951 1435530 75074 95666 64547 1458328 2964507 0 295128 3039582 0 0 0 0 0 0
  8       0 sda 5518 9968 530898 14588 55987 51443 1405112 104687 0 150392 119276 0 0 0 0 0 0
  8       1 sda1 5480 9968 530194 14520 55987 51443 1405112 104687 0 150304 119207 0 0 0 0 0 0
`

func TestDiskRW_Values(t *testing.T) {
	drw := NewDiskRW()
	stats := drw.parseDiskStats(test)
	if len(stats) != 2 {
		t.Errorf("expected len 2, got %d", len(stats))
	}
}
