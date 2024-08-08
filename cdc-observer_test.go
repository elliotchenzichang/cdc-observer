package cdcobserver

import (
	"context"
	"testing"
)

func TestSyncCDCChangeFromDatabase(t *testing.T) {
	opt := &Options{
		DSN:          "127.0.0.1:3307",
		Username:     "elliot_test",
		Password:     "123456",
		EnableDocker: true,
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
