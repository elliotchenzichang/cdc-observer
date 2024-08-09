package cdcobserver

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const ImageName = "mysql"

type DockerClient struct {
	client *client.Client
}

func NewDockerClient() *DockerClient {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	dockerClient := &DockerClient{
		client: cli,
	}
	return dockerClient
}

func (dc *DockerClient) StartMySQLContainer(ctx context.Context) error {
	cli := dc.client
	// todo implement a common function to check if the image had been download
	if !dc.checkIamgeExistence(ctx, ImageName) {
		reader, err := cli.ImagePull(ctx, ImageName, image.PullOptions{})
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
					HostPort: "3307",
				},
			},
		},
	}

	// todo can only create one container and then start it in the future
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: ImageName,
		// todo check why the this is not actually the container name, and also find the approach to set the customized container name
		Hostname: "mysql:cdc-observer:" + RandStringBytesMaskImpr(10),
		Env: []string{
			"MYSQL_ROOT_USERNAME=elliot_test",
			"MYSQL_ROOT_PASSWORD=123456",
			"MYSQL_DATABASE=cdc-observer",
		},
		Tty: false,
	}, hostConfig, nil, nil, "")
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

func (dc *DockerClient) StopAllContianer(ctx context.Context) {
	cli := dc.client

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		log.Fatalf("failed to list containers: %v", err)
	}

	for _, c := range containers {
		if c.State == "running" {
			err := cli.ContainerStop(ctx, c.ID, container.StopOptions{})
			if err != nil {
				log.Printf("failed to stop container %s: %v", c.ID, err)
			} else {
				fmt.Printf("stopped container %s\n", c.ID)
			}
		}
	}
}

func (dc *DockerClient) checkIamgeExistence(ctx context.Context, imageName string) bool {
	cli := dc.client
	_, _, err := cli.ImageInspectWithRaw(ctx, imageName)
	return err == nil
}
