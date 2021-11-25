package coordinator

import (
	"sync"

	"github.com/quarterblue/beehive/internal/job"
	"github.com/quarterblue/beehive/internal/node"
)

type Config struct {
	Name string
}

type Coordinator struct {
	sync.RWMutex
	Config   Config
	NodeList []*node.Node
}

func (c *Coordinator) AddNode(node *node.Node) error {
	c.Lock()
	c.NodeList = append(c.NodeList, node)
	c.Unlock()
	return nil
}

func (c *Coordinator) EditNode() error {
	return nil
}

func (c *Coordinator) RemoveNode() error {
	return nil
}

func (c *Coordinator) AddNodeFromConfig() error {
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
