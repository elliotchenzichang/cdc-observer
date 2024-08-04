package cdcobserver

import (
	"context"

	"github.com/go-mysql-org/go-mysql/canal"
)

type CDCObserver struct {
	river        *canal.Canal
	dockerClient *DockerClient
}

func NewCDCObserver(ctx context.Context, opt *Options) (*CDCObserver, error) {
	if err := opt.validates(); err != nil {
		return nil, err
	}
	observer := &CDCObserver{}
	if opt.EnableDocker {
		dockerClient := NewDockerClient(opt)
		// todo considering add the PGSQL in this repo as well, not urgent, but add a todo here
		dockerClient.StartMySQLContainer(ctx)
		observer.dockerClient = dockerClient
	}
	cfg := canal.NewDefaultConfig()
	cfg.Addr = opt.DSN
	cfg.User = opt.Username
	cfg.Password = opt.Password
	c, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}
	c.SetEventHandler(&CDCObserverHandler{})
	observer.river = c
	return observer, nil
}

func (ob *CDCObserver) Close(ctx context.Context) {
	ob.dockerClient.StopAllContianer(ctx)
	ob.river.Close()
}
