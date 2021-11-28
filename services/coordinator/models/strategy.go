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

type SRoundRobin struct {
	mu    sync.Mutex
	index int
	nodes []*node.Node
}

func (srr *SRoundRobin) Add(node *node.Node) error {
	srr.mu.Lock()
	defer srr.mu.Unlock()
	for _, n := range srr.nodes {
		if n.ID == node.ID {
			return errors.New("node with the same identifier already exists")
		}
	}
	srr.nodes = append(srr.nodes, node)

	return nil
}

func (srr *SRoundRobin) Edit(node *node.Node) error {
	srr.mu.Lock()
	defer srr.mu.Unlock()
	for i, n := range srr.nodes {
		if n.ID == node.ID {
			srr.nodes[i] = node
			return nil
		}
	}

	return errors.New("no such node exists")
}

func (srr *SRoundRobin) Remove(node *node.Node) error {
	srr.mu.Lock()
	defer srr.mu.Unlock()
	for i, n := range srr.nodes {
		if n.ID == node.ID {
			srr.nodes = RemoveIndex(srr.nodes, i)
			return nil
		}
	}

	return errors.New("no such node exists")
}

func (srr *SRoundRobin) Next(job *job.Job) (*node.Node, error) {
	srr.mu.Lock()
	defer srr.mu.Lock()

	if srr.index >= len(srr.nodes) {
		srr.index = 0
	}

	node := &node.Node{
		ID:     srr.nodes[srr.index].ID,
		Name:   srr.nodes[srr.index].Name,
		IpAddr: srr.nodes[srr.index].IpAddr,
		Port:   srr.nodes[srr.index].Port,
	}

	srr.index++

	return node, nil
}

type WRoundRobin struct {
}

func (wrr *WRoundRobin) Add(node *node.Node) error {
	return nil
}

func (wrr *WRoundRobin) Edit(node *node.Node) error {
	return nil
}

func (wrr *WRoundRobin) Remove(node *node.Node) error {
	return nil
}

func (wrr *WRoundRobin) Next(job *job.Job) (*node.Node, error) {
	return nil, nil
}

type LJobs struct {
}

func (lj *LJobs) Add(node *node.Node) error {
	return nil
}

func (lj *LJobs) Edit(node *node.Node) error {
	return nil
}

func (lj *LJobs) Remove(node *node.Node) error {
	return nil
}

func (lj *LJobs) Next(job *job.Job) (*node.Node, error) {
	return nil, nil
}

type CHash struct {
}

func (ch *CHash) Add(node *node.Node) error {
	return nil
}

func (ch *CHash) Edit(node *node.Node) error {
	return nil
}

func (ch *CHash) Remove(node *node.Node) error {
	return nil
}

func (ch *CHash) Next(job *job.Job) (*node.Node, error) {
	return nil, nil
}

func RemoveIndex(s []*node.Node, index int) []*node.Node {
	return append(s[:index], s[index+1:]...)
}
