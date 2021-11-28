package node

import (
	"github.com/google/uuid"
)

// Represents the node specification
type Node struct {
	ID       uuid.UUID
	Name     string
	IpAddr   string
	Port     string
	JobCount uint32
	Spec     *Spec
}

// The machine specification
type Spec struct {
	CPUmodel   string
	CPUmhz     float64
	CPUcore    uint64
	Memory     uint64
	MemoryFree uint64
	Disk       uint64
	DiskFree   uint64
	OS         string
	Platform   string
	KernelArch string
	Hostname   string
	Uptime     uint64
	BootTime   uint64
}

// Create a new node by fetching machine specfication information
func NewNode(name, ipaddr, port string, spec *Spec) *Node {

	node := &Node{
		Name:     name,
		IpAddr:   ipaddr,
		Port:     port,
		JobCount: 0,
		Spec:     spec,
	}

	return node
}

func NewSpec(mem, memalloc, disk, diskalloc uint64, os, plat, host string) *Spec {
	return &Spec{
		Memory:     mem,
		MemoryFree: memalloc,
		Disk:       disk,
		DiskFree:   diskalloc,
		OS:         os,
		Platform:   plat,
		Hostname:   host,
	}
}
