package coordinator

import (
	"errors"
	"sync"

	"github.com/quarterblue/beehive/internal/job"
	"github.com/quarterblue/beehive/internal/node"
)

type Manager interface {
	Add(*node.Node) error
	Edit(*node.Node) error
	Remove(*node.Node) error
	Next(*job.Job) (*node.Node, error)
}

type WeightedRR struct {
}

func (nm *WeightedRR) Add(node *node.Node) error {
	return nil
}

func (nm *WeightedRR) Edit(node *node.Node) error {
	return nil
}

func (nm *WeightedRR) Remove(node *node.Node) error {
	return nil
}

func (nm *WeightedRR) Next(job *job.Job) (*node.Node, error) {
	return nil, nil
}

type SRoundRobin struct {
	mu    sync.Mutex
	index uint64
	nodes []*node.Node
}

func (nm *SRoundRobin) Add(node *node.Node) error {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	for _, n := range nm.nodes {
		if n.ID == node.ID {
			return errors.New("node with the same identifier already exists")
		}
	}
	nm.nodes = append(nm.nodes, node)

	return nil
}

func (nm *SRoundRobin) Edit(node *node.Node) error {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	for i, n := range nm.nodes {
		if n.ID == node.ID {
			nm.nodes[i] = node
			return nil
		}
	}

	return errors.New("no such node exists")
}

func (nm *SRoundRobin) Remove(node *node.Node) error {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	for i, n := range nm.nodes {
		if n.ID == node.ID {
			nm.nodes = RemoveIndex(nm.nodes, i)
			return nil
		}
	}

	return errors.New("no such node exists")
}

func (nm *SRoundRobin) Next(job *job.Job) (*node.Node, error) {
	return nil, nil
}

func RemoveIndex(s []*node.Node, index int) []*node.Node {
	return append(s[:index], s[index+1:]...)
}
