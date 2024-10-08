package cdcobserver

import (
	"context"
	"testing"
)

func TestStartMySQLContainer(t *testing.T) {
	dockerClient, _ := NewDockerClient()
	ctx := context.Background()
	_ = dockerClient.StartMySQLContainer(ctx)
}
