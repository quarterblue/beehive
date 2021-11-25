package api

// import (
// 	"sync"

// 	"github.com/manucorporat/stats"
// )

// var (
// 	ips        = stats.New()
// 	messages   = stats.New()
// 	users      = stats.New()
// 	mutexStats sync.RWMutex
// 	savedStats map[string]uint64
// )

// func connectedUsers() uint64 {
// 	connected := users.Get("connected") - users.Get("disconnected")
// 	if connected < 0 {
// 		return 0
// 	}
// 	return uint64(connected)
// }

// // Stats returns savedStats data.
// func Stats() map[string]uint64 {
// 	mutexStats.RLock()
// 	defer mutexStats.RUnlock()

// 	return savedStats
// }
