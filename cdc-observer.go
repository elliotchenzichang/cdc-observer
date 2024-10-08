package cdcobserver

import (
	dockerapi "cdc-observer/docker_api"
	"context"
	"errors"
	"time"

	"cdc-observer/database"

	"github.com/go-mysql-org/go-mysql/canal"
)

type CDCObserver struct {
	enableDocker  bool
	containername string
	containerPort int
	username      string
	password      string
	addr          string
	dbName        string
	river         *canal.Canal
	dockerClient  *dockerapi.DockerClient
	// only one database is enough for the goal of this project
	db *database.Database
}

func NewCDCObserver(opt *Options) (*CDCObserver, error) {
	if err := opt.validates(); err != nil {
		return nil, err
	}
	observer := &CDCObserver{}
	observer.dbName = opt.DatabaseName
	if opt.EnableDocker {
		observer.enableDocker = true
		observer.containername = opt.ContainerName
		dockerClient, err := dockerapi.NewDockerClient()
		if err != nil {
			return nil, err
		}
		observer.dockerClient = dockerClient
	}
	observer.containerPort = opt.ContainerPort
	observer.username = opt.Username
	observer.password = opt.Password
	observer.addr = opt.DSN
	return observer, nil
}

func (ob *CDCObserver) Start(ctx context.Context) error {
	if ob.enableDocker {
		// todo considering add the PGSQL in this repo as well, not urgent, but add a todo here
		go ob.dockerClient.StartMySQLContainer(ctx)
		time.Sleep(3 * time.Second)

	}
	cfg := canal.NewDefaultConfig()
	cfg.Addr = ob.addr
	cfg.User = ob.username
	cfg.Password = ob.password
	c, err := canal.NewCanal(cfg)
	if err != nil {
		return err
	}

	if c == nil {
		return errors.New("the river is empty, please check if your database enable the binlog and the log style is ROW")
	}
	c.SetEventHandler(&CDCObserverHandler{})

	ob.river = c
	return nil
}

func (ob *CDCObserver) Close(ctx context.Context) error {
	if ob.enableDocker {
		ob.dockerClient.StopAllContainers(ctx)
	}
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
