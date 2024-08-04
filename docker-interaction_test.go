package cdcobserver

import (
	"context"
	"testing"
)

func TestStartMySQLContainer(t *testing.T) {
	dockerClient := DockerClient{}
	ctx := context.Background()
	dockerClient.StartMySQLContainer(ctx)
}
