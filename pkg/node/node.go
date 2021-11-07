package node

type Node struct {
	ID              string
	Name            string
	IpAddr          string
	Port            string
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	JobCount        int
}
