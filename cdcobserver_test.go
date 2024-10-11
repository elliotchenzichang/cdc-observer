package cdcobserver

import (
	"context"
	"testing"
	"time"
)

func TestSyncCDCChangeFromDatabase(t *testing.T) {
	ctx := context.Background()
	cdcObserver, err := NewCDCObserver()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		cdcObserver.dc.RemoveAllContainers(ctx)
	}()

	ch := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		cdcObserver.Close(ctx)
		close(ch)
	}()

	err = cdcObserver.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("start to listen the change event from local MySQL")
	<-ch
}
