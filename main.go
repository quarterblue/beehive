package main

import (
	// "archive/tar"
	// "bytes"
	// "context"
	// "fmt"
	// "io"
	// "io/ioutil"
	// "log"
	// "os"

	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/client"
	"fmt"

	"github.com/quarterblue/beehive/internal/node"
)

func main() {

	node := node.NewNode("name", "ip", "port")
	fmt.Println(node)
	// fmt.Println("Welcome to beehive Job Scheduler")
	// ctx := context.Background()

	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	// if err != nil {
	// 	panic(err)
	// }

	// containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// for _, container := range containers {
	// 	fmt.Println(container.Names)
	// }

	// buf := new(bytes.Buffer)
	// tw := tar.NewWriter(buf)
	// defer tw.Close()

	// dockerFile := "dockerfile"

	// dockerFileReader, err := os.Open("dockerfile")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	// if err != nil {
	// 	log.Fatal(err, " :unable to read dockerfile")
	// }

	// tarHeader := &tar.Header{
	// 	Name: dockerFile,
	// 	Size: int64(len(readDockerFile)),
	// }
	// err = tw.WriteHeader(tarHeader)
	// if err != nil {
	// 	log.Fatal(err, " :unable to write tar header")
	// }
	// _, err = tw.Write(readDockerFile)
	// if err != nil {
	// 	log.Fatal(err, " :unable to write tar body")
	// }
	// dockerFileTarReader := bytes.NewReader(buf.Bytes())

	// imageBuildResponse, err := cli.ImageBuild(
	// 	ctx,
	// 	dockerFileTarReader,
	// 	types.ImageBuildOptions{
	// 		Context:    dockerFileTarReader,
	// 		Dockerfile: dockerFile,
	// 		Remove:     true})

	// if err != nil {
	// 	log.Fatal(err, " :unable to build docker image")
	// }

	// defer imageBuildResponse.Body.Close()

	// _, err = io.Copy(os.Stdout, imageBuildResponse.Body)

	// if err != nil {
	// 	log.Fatal(err, " :unable to read image build response")
	// }

}
