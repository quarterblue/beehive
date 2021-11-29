package coordinator

import (
	"log"
	"sync"
	"testing"

	"github.com/quarterblue/beehive/internal/node"
)

func TestSRoundRobinAdd(t *testing.T) {
	srr := &SRoundRobin{
		mu:    sync.Mutex{},
		index: 0,
		nodes: make([]*node.Node, 0),
	}

	nodeOne := &node.Node{
		ID:       2,
		Name:     "nodeone",
		IpAddr:   "171.28.49.10",
		Port:     "3001",
		JobCount: 0,
	}

	nodeTwo := &node.Node{
		ID:       3,
		Name:     "nodetwo",
		IpAddr:   "171.28.50.10",
		Port:     "3002",
		JobCount: 0,
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

func TestLJobsAdd(t *testing.T) {
	l := &LJobs{
		mu:    sync.Mutex{},
		nodes: make([]*node.Node, 0),
	}

	nodeOne := &node.Node{
		ID:       1,
		Name:     "testnodeOne",
		IpAddr:   "182.10.2.30",
		Port:     "3001",
		JobCount: 7,
	}

	nodeTwo := &node.Node{
		ID:       2,
		Name:     "testnodeTwo",
		IpAddr:   "182.10.2.31",
		Port:     "3001",
		JobCount: 2,
	}

	nodeThree := &node.Node{
		ID:       3,
		Name:     "testnodeThree",
		IpAddr:   "182.10.2.32",
		Port:     "3001",
		JobCount: 6,
	}

	nodeFour := &node.Node{
		ID:       4,
		Name:     "testnodeFour",
		IpAddr:   "182.10.2.33",
		Port:     "3001",
		JobCount: 0,
	}

	l.Add(nodeOne)
	l.Add(nodeTwo)
	l.Add(nodeThree)
	l.Add(nodeFour)
	for i, v := range l.nodes {
		log.Println(i, v)
	}

}

func TestLJobsEdit(t *testing.T) {

}

func TestLJobsRemove(t *testing.T) {

}
func TestLJobsNext(t *testing.T) {

}
