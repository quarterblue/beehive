package coordinator

import (
	"sync"

	"github.com/quarterblue/beehive/internal/job"
	"github.com/quarterblue/beehive/internal/node"
)

type Strategy int

const (
	WeightedRoundRobin Strategy = iota
	RoundRobin
	LeastJobs
	ConsistentHashing
)

type Config struct {
	Name string
}

// Coordinator
type Coordinator struct {
	mu          sync.RWMutex
	config      Config
	nodeManager Manager
}

func NewNodeManager(strategy Strategy) *Manager {
	return nil
}

func NewCoordinator(cfg Config) *Coordinator {
	return &Coordinator{
		mu:          sync.RWMutex{},
		config:      cfg,
		nodeManager: nil,
	}
}

func (c *Coordinator) AddNode(node *node.Node) error {
	c.mu.Lock()
	c.nodeManager.Add(node)
	c.mu.Unlock()
	return nil
}

func (c *Coordinator) EditNode() error {
	return nil
}

func (c *Coordinator) RemoveNode() error {
	return nil
}

func (c *Coordinator) AddNodesFromConfig() error {
	return nil
}

type result struct {
	msg string
	err error
}

func ScheduleJob(job job.DockerJob) result {

	return result{
		msg: "",
		err: nil,
	}
}
