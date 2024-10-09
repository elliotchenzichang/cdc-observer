package dockerapi

import (
	"cdc-observer/constant"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartMySQLContainer(t *testing.T) {
	dockerClient, err := NewDockerClient()
	assert.NoError(t, err, "Failed to create Docker client")

	ctx := context.Background()
	err = dockerClient.StartMySQLContainer(ctx)
	assert.NoError(t, err, "Failed to start MySQL container")

	defer func() {
		dockerClient.StopAllContainers(ctx)
		dockerClient.RemoveAllContainers(ctx)
	}()

	containerName := dockerClient.ContainerName(constant.MysqlImageName)
	cj, err := dockerClient.ContainerInfo(ctx, containerName)
	assert.NoError(t, err, "Failed to get MySQL container info")
	assert.True(t, cj.State.Running, "MySQL container should be running")

	dockerClient.StopAllContainers(ctx)
	cj, err = dockerClient.ContainerInfo(ctx, containerName)
	assert.NoError(t, err, "Failed to get MySQL container info")
	assert.False(t, cj.State.Running, "MySQL container should be stopped")
}
