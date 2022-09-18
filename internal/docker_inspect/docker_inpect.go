package dockerinspect

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)


type DockerInspector struct{

}

func (di *DockerInspector) Inspect 

func InitDockerClient() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
		stat, _ := cli.ContainerInspect(context.Background(), container.ID)
		//fmt.Println(cli.ContainerInspect(context.Background(), container.ID))
		//var data map[string]interface{}

		fmt.Println(stat.State.Status)
	}

}
