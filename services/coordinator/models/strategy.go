package coordinator

import (
	"errors"
	"sort"
	"sync"

	"github.com/quarterblue/beehive/internal/node"
)

var (
	ErrNoNode   error = errors.New("no such node exists")
	ErrSameId   error = errors.New("node with same identifier already exists")
	ErrEmptyJob error = errors.New("empty job list")
)

type Manager interface {
	Add(*node.Node) error
	Edit(*node.Node) error
	Remove(id int64) error
	Next() (*node.Node, error)
}

// Implementation of Simple Round Robin load balancer
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
			return ErrSameId
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

	return ErrNoNode
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

	return ErrNoNode
}

func (srr *SRoundRobin) Next() (*node.Node, error) {
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

// Implementation of Weighted (machine performance) Round Robin load balancer
type WRoundRobin struct {
	high     []*node.Node
	medium   []*node.Node
	low      []*node.Node
	specList []node.LoadSpec
}

func NewWRR(lSpec, mSpec, hSpec node.LoadSpec) *WRoundRobin {
	return &WRoundRobin{
		high:     make([]*node.Node, 0),
		medium:   make([]*node.Node, 0),
		low:      make([]*node.Node, 0),
		specList: []node.LoadSpec{lSpec, mSpec, hSpec},
	}
}

func (wrr *WRoundRobin) Add(node *node.Node) error {
	// for spec := range wrr.specList {
	// }
	return nil
}

func (wrr *WRoundRobin) Edit(node *node.Node) error {
	return nil
}

func (wrr *WRoundRobin) Remove(id int64) error {
	return nil
}

func (wrr *WRoundRobin) Next() (*node.Node, error) {
	return nil, nil
}

// Implements least jobs load balancer
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
	// Keep the list always sorted in descending order.
	sort.Sort(sort.Reverse(l))
	return nil
}

func (l *LJobs) Edit(node *node.Node) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, n := range l.nodes {
		if n.ID == node.ID {
			l.nodes[i] = node
			break
		}
	}
	// Keep the list always sorted in descending order.
	sort.Sort(sort.Reverse(l))
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

	return ErrNoNode
}

// Next returns the node with least jobs, which we pop from end of the slice
// Since the slice is always sorted in descending job order
func (l *LJobs) Next() (*node.Node, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	var i int
	if i = len(l.nodes) - 1; i <= -1 {
		return nil, errors.New("empty job list")
	}
	x := l.nodes[i]
	l.nodes = append(l.nodes[:i], l.nodes[i+1:]...)
	return x, nil
}

// Implements consistent hashing with bounded loads, load balancer
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

func (ch *CHash) Next() (*node.Node, error) {
	return nil, nil
}

// Util functions

func RemoveIndex(s []*node.Node, index int) []*node.Node {
	return append(s[:index], s[index+1:]...)
}
