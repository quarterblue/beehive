package node

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type Node struct {
	ID              uuid.UUID
	Name            string
	IpAddr          string
	Port            string
	Memory          uint64
	MemoryAllocated uint64
	Disk            uint64
	DiskAllocated   uint64
	JobCount        uint32
}

func NewNode(name, ipaddr, port string) *Node {

	node := &Node{
		Name:   name,
		IpAddr: ipaddr,
		Port:   port,
	}

	hInfo, _ := host.Info()
	// cpuStat, _ := cpu.Info()
	// diskStat, _ := disk.Usage("/")
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	return node
}
