package coordinator

type Strategy int

const (
	RoundRobin Strategy = iota
	LeastConnection
	LeastResponseTime
	ConsistentHash
)

type Scheduler struct {
	LoadBalancer Strategy
}

func NewScheduler(strategy Strategy) {

}

func StrategyRR() {

}

func StrategyLC() {

}

func StrategyLR() {

}

func StrategyCH() {

}
