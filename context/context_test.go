package context

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctxOriginal, cancelOriginal := context.WithTimeoutCause(context.Background(), 3*time.Minute, errors.New("original cause"))
	ctx, cancel := context.WithTimeoutCause(ctxOriginal, 1*time.Minute, errors.New("child cause"))
	defer cancel()
	defer cancelOriginal()

	time.Sleep(50 * time.Millisecond)
	fmt.Printf("parent context: %+v\n", ctxOriginal)
	fmt.Printf("child context: %+v\n", ctx)

	select {
	case <-time.After(10 * time.Millisecond):
		fmt.Println("feliz")
	case <-ctx.Done():
		fmt.Println("child context", ctx.Err())
	case <-ctxOriginal.Done():
		fmt.Println("parent context", ctxOriginal.Err())
	}
	fmt.Println("Hello, 世界")
}
