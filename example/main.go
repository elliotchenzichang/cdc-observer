package main

import (
	dockerapi "cdc-observer/docker_api"
	"context"
)

func main() {
	dockerClient, _ := dockerapi.NewDockerClient()
	ctx := context.Background()
	_ = dockerClient.StartMySQLContainer(ctx)
}
