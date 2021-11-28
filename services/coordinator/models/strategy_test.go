package coordinator

import (
	"sync"
	"testing"

	"github.com/quarterblue/beehive/internal/node"
)

func TestSRoundRobinAdd(t *testing.T) {
	srr := &SRoundRobin{
		mu:    sync.Mutex{},
		index: 0,
		nodes: make([]*node.Node, 10),
	}

	id := AddrToId("171.28.49.10", "3001")

	nodeOne := &node.Node{
		ID:       id,
		Name:     "nodeone",
		IpAddr:   "171.28.49.10",
		Port:     "3001",
		JobCount: 0,
		Spec:     nil,
	}

	idtwo := AddrToId("171.28.50.10", "3002")

	nodeTwo := &node.Node{
		ID:       idtwo,
		Name:     "nodetwo",
		IpAddr:   "171.28.50.10",
		Port:     "3002",
		JobCount: 0,
		Spec:     nil,
	}

	srr.Add(nodeOne)
	srr.Add(nodeTwo)

	if len(srr.nodes) != 2 {
		t.Errorf("got %d, want %d", len(srr.nodes), 2)
	}
}

func TestSRoundRobinEdit(t *testing.T) {

}

func TestSRoundRobinRemove(t *testing.T) {

}
func TestSRoundRobinNext(t *testing.T) {

}
