package cdcobserver

import (
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
	dockerClient  *DockerClient
	// only one database is enough for the goal of this project
	db *database.Database
}

func NewCDCObserver(ctx context.Context, opt *Options) (*CDCObserver, error) {
	if err := opt.validates(); err != nil {
		return nil, err
	}
	observer := &CDCObserver{}
	observer.dbName = opt.DatabaseName
	if opt.EnableDocker {
		observer.enableDocker = true
		observer.containername = opt.ContainerName
		dockerClient := NewDockerClient(observer.dbName)
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
		ob.dockerClient.StopAllContianer(ctx)
	}
	ob.river.Close()
	return nil
}
