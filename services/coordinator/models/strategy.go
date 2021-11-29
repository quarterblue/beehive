package coordinator

import (
	"errors"
	"sort"
	"sync"

	"github.com/quarterblue/beehive/internal/job"
	"github.com/quarterblue/beehive/internal/node"
)

type Manager interface {
	Add(*node.Node) error
	Edit(*node.Node) error
	Remove(id int64) error
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

func (srr *SRoundRobin) Remove(id int64) error {
	srr.mu.Lock()
	defer srr.mu.Unlock()
	for i, n := range srr.nodes {
		if n.ID == id {
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

type Quality int

const (
	qHigh Quality = iota
	qMedium
	qLow
)

type LoadSpec struct {
	CPUMhz  float64
	Memory  uint64
	Disk    uint64
	Quality Quality
}
type WRoundRobin struct {
	high   []*node.Node
	medium []*node.Node
	low    []*node.Node
	hSpec  LoadSpec
	mSpec  LoadSpec
	lSpec  LoadSpec
}

func (wrr *WRoundRobin) Add(node *node.Node) error {
	return nil
}

func (wrr *WRoundRobin) Edit(node *node.Node) error {
	return nil
}

func (wrr *WRoundRobin) Remove(id int64) error {
	return nil
}

func (wrr *WRoundRobin) Next(job *job.Job) (*node.Node, error) {
	return nil, nil
}

type LJobs struct {
	mu    sync.Mutex
	nodes []*node.Node
}

// Implements sort.Interface Len for custom sorting
func (l *LJobs) Len() int {
	return len(l.nodes)
}

// Implement sort.Interface Less for custom sorting
func (l *LJobs) Less(i, j int) bool {
	return l.nodes[i].JobCount < l.nodes[j].JobCount
}

// Implement sort.Interface Swap for custom sorting
func (l *LJobs) Swap(i, j int) {
	l.nodes[i], l.nodes[j] = l.nodes[j], l.nodes[i]
}

func (l *LJobs) Add(node *node.Node) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.nodes = append(l.nodes, node)
	// Keep the list always sorted
	sort.Sort(l)
	return nil
}

func (l *LJobs) Edit(node *node.Node) error {
	return nil
}

func (l *LJobs) Remove(id int64) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, n := range l.nodes {
		if n.ID == id {
			l.nodes = RemoveIndex(l.nodes, i)
			return nil
		}
	}

	return errors.New("no such node exists")
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

func (ch *CHash) Remove(id int64) error {
	return nil
}

func (ch *CHash) Next(job *job.Job) (*node.Node, error) {
	return nil, nil
}

func RemoveIndex(s []*node.Node, index int) []*node.Node {
	return append(s[:index], s[index+1:]...)
}
