package api

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartGin(port string) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	// router.Use(rateLimit, gin.Recovery())
	router.GET("/", index)
	router.GET("/job", jobGETALL)
	router.GET("/job/:jobid", jobGETONE)
	router.POST("/job", jobCREATE)
	router.DELETE("/job/:jobid", jobDELETE)

	router.Run(port)
}

func main() {
	ConfigRuntime()
	StartGin(":9001")
}
