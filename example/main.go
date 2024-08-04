package main

import (
	cdc "cdc-observer"
	"context"
)

func main() {
	opt := &cdc.Options{}
	dockerClient := cdc.NewDockerClient(opt)
	ctx := context.Background()
	dockerClient.StartMySQLContainer(ctx)
}
