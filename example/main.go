package main

import (
	cdc "cdc-observer"
	"context"
)

func main() {
	dockerClient := cdc.NewDockerClient("elliot_test_name")
	ctx := context.Background()
	dockerClient.StartMySQLContainer(ctx)
}
