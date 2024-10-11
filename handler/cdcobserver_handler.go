package handler

import (
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/siddontang/go-log/log"
)

type CDCObserverHandler struct {
	canal.DummyEventHandler
}

// todo build a tool for print the CDC event in terminal in table style
func (h *CDCObserverHandler) OnRow(e *canal.RowsEvent) error {
	log.Infof("%s %v\n", e.Action, e.Rows)
	return nil
}

func (h *CDCObserverHandler) String() string {
	return "CDCObserverHandler"
}
