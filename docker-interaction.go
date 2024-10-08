package cdcobserver

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// the prefix of the name of the cdc-observer containers, like /cdc-observer-mysql, /cdc-observer-pgsql
const ContainerNamePrefix = "cdc-observer-"

// the image name of the database instance in the container
const (
	MysqlImageName = "mysql"
)

// all the databases share this settings
const (
	DatabaseName     = "cdc-observer"
	DatabaseUsername = "root"
	DatabasePassword = "cdc-observer-password"
)

type DockerClient struct {
	client *client.Client
}

func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	dockerClient := &DockerClient{
		client: cli,
	}
	return dockerClient, nil
}

func (dc *DockerClient) StartMySQLContainer(ctx context.Context) error {
	cli := dc.client
	// todo implement a common function to check if the image had been download
	if !dc.checkImageExistence(ctx, MysqlImageName) {
		reader, err := cli.ImagePull(ctx, MysqlImageName, image.PullOptions{})
		if err != nil {
			return err
		}
		defer reader.Close()
		// cli.ImagePull is asynchronous.
		// The reader needs to be read completely for the pull operation to complete.
		// If stdout is not required, consider using io.Discard instead of os.Stdout.
		io.Copy(os.Stdout, reader)
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"3306/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "0", // use 0 to let docker automatically choose a free port
				},
			},
		},
	}

	containerName := fmt.Sprintf("/%s%s", ContainerNamePrefix, MysqlImageName)

	// Check if the container named /cdc-observer-mysql already exists
	if exists, err := dc.checkContainerExistence(ctx, containerName); err != nil {
		return err
	} else if exists {
		if err := dc.handleExistingContainer(ctx, containerName); err != nil {
			return err
		}
		return nil
	}

	// If the container doesn't exist, create it
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: MysqlImageName,
		Env: []string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", DatabasePassword),
			fmt.Sprintf("MYSQL_DATABASE=%s", DatabaseName),
		},
		Tty: false,
	}, hostConfig, nil, nil, containerName)
	if err != nil {
		return err
	}
	log.Printf("successfully create a new MySQL container, containerID: %s, warnings: %+v", resp.ID, resp.Warnings)

	// todo check how to print the docker running log
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return err
	}

	log.Printf("successfully start a new mysql SQL container with ID: %s", resp.ID)

	_, err = cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	// stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return nil
}

func (dc *DockerClient) StopAllContainers(ctx context.Context) {
	containers, err := dc.containers(ctx)
	if err != nil {
		log.Printf("failed to list containers: %v", err)
	}

	for _, c := range containers {
		if c.State == "running" {
			err := dc.client.ContainerStop(ctx, c.ID, container.StopOptions{})
			if err != nil {
				log.Printf("failed to stop container %s: %v", c.ID, err)
			} else {
				log.Printf("stopped container %s\n", c.ID)
			}
		}
	}
}

func (dc *DockerClient) checkImageExistence(ctx context.Context, imageName string) bool {
	_, _, err := dc.client.ImageInspectWithRaw(ctx, imageName)
	return err == nil
}

func (dc *DockerClient) checkContainerExistence(ctx context.Context, containerName string) (bool, error) {
	if _, err := dc.client.ContainerInspect(ctx, containerName); err != nil {
		if client.IsErrNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (dc *DockerClient) containers(ctx context.Context) ([]types.Container, error) {
	return dc.client.ContainerList(ctx, container.ListOptions{All: true})
}

func (dc *DockerClient) handleExistingContainer(ctx context.Context, containerName string) error {
	// Check the container's state
	cj, err := dc.client.ContainerInspect(ctx, containerName)
	if err != nil {
		return err
	}

	switch cj.State.Status {
	case "running":
		// If the container is running, do nothing
		return nil
	case "exited", "created", "paused":
		// If the container exists but is not running, start it
		if err := dc.client.ContainerStart(ctx, containerName, container.StartOptions{}); err != nil {
			return err
		}
		return nil
	case "restarting":
		// Wait for the container to finish restarting
		for cj.State.Status == "restarting" {
			time.Sleep(100 * time.Millisecond)
			cj, err = dc.client.ContainerInspect(ctx, containerName)
			if err != nil {
				return err
			}
		}
		if cj.State.Status != "running" {
			if err := dc.client.ContainerStart(ctx, containerName, container.StartOptions{}); err != nil {
				return err
			}
		}
		return nil
	default:
		// For any other state, attempt to start the container
		if err := dc.client.ContainerStart(ctx, containerName, container.StartOptions{}); err != nil {
			return err
		}
		return nil
	}
}
