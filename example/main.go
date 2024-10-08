package main

import (
	cdc "cdc-observer"
	"context"
)

func main() {
	dockerClient, _ := cdc.NewDockerClient()
	ctx := context.Background()
	_ = dockerClient.StartMySQLContainer(ctx)
}
