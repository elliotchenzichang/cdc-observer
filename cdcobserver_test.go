package cdcobserver

import (
	"context"
	"testing"
)

func TestSyncCDCChangeFromDatabase(t *testing.T) {
	ctx := context.Background()
	cdcObserver, err := NewCDCObserver()
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
