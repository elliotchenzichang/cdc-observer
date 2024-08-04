package cdcobserver

import "github.com/go-mysql-org/go-mysql/canal"

type CDCObserver struct {
	river *canal.Canal
}

func NewCDCObserver(opt *Options) *CDCObserver {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = opt.DSN
	cfg.User = opt.Username
	cfg.Password = opt.Password
	c, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}
	c.SetEventHandler(&CDCObserverHandler{})
	observer := &CDCObserver{
		river: c,
	}
	return observer
}
