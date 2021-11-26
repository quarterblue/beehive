package coordinator

type Strategy int

const (
	WeightedRoundRobin Strategy = iota
	LeastUtilized
	LeastResponseTime
	ConsistentHash
)

type Scheduler struct {
	LoadBalancer Strategy
}

func NewScheduler(strategy Strategy) {

}

func StrategyWRR() {

}

func StrategyLU() {

}

func StrategyLR() {

}

func StrategyCH() {

}
