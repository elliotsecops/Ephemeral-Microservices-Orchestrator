package container

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func CreateContainer(image string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: image,
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("Container created:", resp.ID)
}

func RemoveContainer(containerID string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	err = cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		panic(err)
	}

	fmt.Println("Container destroyed:", containerID)
}
