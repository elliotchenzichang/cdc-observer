package cdcobserver

import (
	"cdc-observer/constant"
	dockerapi "cdc-observer/docker_api"
	"cdc-observer/handler"
	"context"
	"errors"
	"fmt"

	"cdc-observer/database"

	"github.com/go-mysql-org/go-mysql/canal"
)

type CDCObserver struct {
	river *canal.Canal
	dc    *dockerapi.DockerClient
	db    *database.Database // Only support MySQL for now
}

func NewCDCObserver() (*CDCObserver, error) {
	observer := &CDCObserver{}
	dockerClient, err := dockerapi.NewDockerClient()
	if err != nil {
		return nil, err
	}
	observer.dc = dockerClient
	return observer, nil
}

func (ob *CDCObserver) Start(ctx context.Context) error {
	// Start the MySQL container
	if err := ob.dc.StartMySQLContainer(ctx); err != nil {
		return err
	}

	// Get the container name and port
	containerName := ob.dc.ContainerName(constant.MysqlImageName)
	port, err := ob.dc.ContainerPort(ctx, containerName)
	if err != nil {
		return err
	}

	// Create a new database connection
	db, err := database.NewDatabase(port)
	if err != nil {
		return err
	}
	ob.db = db

	// Configure and create a new Canal instance
	cfg := canal.NewDefaultConfig()
	cfg.Addr = fmt.Sprintf("%s:%s", constant.DatabaseHost, port)
	cfg.User = constant.DatabaseUsername
	cfg.Password = constant.DatabasePassword
	// Disable dump by setting Dump.ExecutionPath to empty string
	cfg.Dump.ExecutionPath = ""
	// Exclude all databases except cdc-observer
	cfg.ExcludeTableRegex = []string{"[^cdc-observer]\\..*"}
	cfg.IncludeTableRegex = []string{"cdc-observer\\..*"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		return err
	}

	// Check if the Canal instance was created successfully
	if c == nil {
		return errors.New("the river is empty, please check if your database enable the binlog and the log style is ROW")
	}

	// Set the event handler and start the Canal
	c.SetEventHandler(&handler.CDCObserverHandler{})
	ob.river = c
	if err := ob.river.Run(); err != nil {
		return err
	}

	return nil
}

func (ob *CDCObserver) Close(ctx context.Context) error {
	ob.dc.StopAllContainers(ctx)
	ob.river.Close()
	return nil
}

func (ob *CDCObserver) AddTable(name string, table *database.Table) error {
	return ob.db.AddTable(table)
}

func (ob *CDCObserver) DeleteTable(name string) error {
	return ob.db.DeleteTable(name)
}

func (ob *CDCObserver) ApplyDB() error {
	return ob.db.Apply()
}

func (ob *CDCObserver) Clean() error {
	return ob.db.Clean()
}
