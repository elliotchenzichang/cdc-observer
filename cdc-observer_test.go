package cdcobserver

import (
	"context"
	"testing"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/siddontang/go/log"
)

func TestSyncCDCChangeFromDatabase(t *testing.T) {
	opt := &Options{
		DSN:          "127.0.0.1:3307",
		Username:     "root",
		Password:     "123456",
		EnableDocker: false,
		DatabaseName: "elliot_test_database",
	}
	ctx := context.Background()
	cdcObserver, err := NewCDCObserver(ctx, opt)
	if err != nil {
		t.Fatal(err)
	}
	err = cdcObserver.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("start to listen the change event from local MySQL")

	err = cdcObserver.Close(ctx)
	if err != nil {
		t.Fatal(err)
	}

}

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	log.Infof("%s %v\n", e.Action, e.Rows)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func TestCanal(t *testing.T) {

	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3307"
	cfg.User = "root"
	cfg.Password = "123456"
	// We only care table canal_test in test db
	// cfg.Dump.TableDB = "cdc-observer"
	// cfg.Dump.Tables = []string{"canal_test"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	c.Run()
}
