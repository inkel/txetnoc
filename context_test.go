package txetnoc

import (
	"context"
	"testing"
	"time"
)

func TestWithChildren(t *testing.T) {
	cancelCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//cancel()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx := WithChildren(cancelCtx, timeoutCtx)

	<-ctx.Done()

	t.Logf("ERR %p %#[1]v", ctx.Err())
	t.Logf("CERR %p %[1]v", cancelCtx.Err())
	t.Logf("TERR %p %[1]v", timeoutCtx.Err())
}
