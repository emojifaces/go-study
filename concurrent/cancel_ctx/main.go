package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func f1(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("f1 canceled")
	case <-time.After(5 * time.Second):
		fmt.Println("f1 done")
	}
}

func f2() error {
	time.Sleep(3 * time.Second)

	return errors.New("unknown error")
}

// 发出取消事件
func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		f1(ctx)
	}()

	if err := f2(); err != nil {
		cancel()
	}

}
