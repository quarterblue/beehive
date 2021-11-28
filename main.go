package main

import ( // "archive/tar"
	// "bytes"
	// "context"
	// "fmt"
	// "io"
	// "io/ioutil"
	// "log"
	// "os"
	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/client"

	"context"
	"fmt"
	"log"

	"github.com/quarterblue/beehive/services/coordinator/api"
	"github.com/quarterblue/beehive/services/worker"
	"github.com/quarterblue/beehive/services/worker/pb"
)

func main() {

	worker := worker.NewWorker("123", "testpc")

	ctx := context.Background()
	pbr, err := worker.MachineSpec(ctx, &pb.SpecRequest{})

	if err != nil {
		log.Println(err)
	}
	// node := node.NewNode("name", "ip", "port")
	// fmt.Println(node)

	fmt.Println(pbr)

	api.Server()
}
