package main

import (
	cdc "cdc-observer"
	"context"
)

func main() {
	dockerClient := cdc.NewDockerClient()
	ctx := context.Background()
	dockerClient.StartMySQLContainer(ctx)
}
