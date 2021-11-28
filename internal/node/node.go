package node

// Represents the node specification
type Node struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IpAddr   string `json:"ipadrr"`
	Port     string `json:"port"`
	JobCount uint32 `json:"jobcount"`
	// Spec     *Spec  `json:"spec"`
}

// The machine specification
// This information is retrieved when a node is first initialized
// We use this information to dynamically load balance jobs to appropriate nodes
type Spec struct {
	CPUmodel   string  `json:"cpu_model"`
	CPUmhz     float64 `json:"cpu_mhz"`
	CPUcore    uint64  `json:"cpu_core"`
	Memory     uint64  `json:"memory"`
	MemoryFree uint64  `json:"memory_free"`
	Disk       uint64  `json:"disk"`
	DiskFree   uint64  `json:"disk_free"`
	OS         string  `json:"os"`
	Platform   string  `json:"platform"`
	KernelArch string  `json:"kernel_arch"`
	Hostname   string  `json:"hostname"`
	Uptime     uint64  `json:"uptime"`
	BootTime   uint64  `json:"boot_time"`
}

// Create a new node by fetching machine specfication information
func NewNode(id int64, name, ipaddr, port string, spec *Spec) *Node {

	node := &Node{
		ID:       id,
		Name:     name,
		IpAddr:   ipaddr,
		Port:     port,
		JobCount: 0,
		// Spec:     spec,
	}

	return node
}

// Converts IPv4 address and port to ipaddr:port format
func AddrToID(addr, port string) string {
	return addr + ":" + port
}
