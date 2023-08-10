package monitors

import "mon2http/src/sources"

const diskRWType = "disk_rw"

func init() {
	registerMonitor(NewDiskRW(sources.NewDiskRW()))
}

type DiskRW struct {
	BaseMonitor
}

func NewDiskRW(s *sources.DiskRW) *DiskRW {
	m := DiskRW{}
	m.source = s
	m.typeString = diskRWType
	return &m
}
