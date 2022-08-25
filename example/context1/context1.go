package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond //?1ms

func main() {
	//*创建截止时间
	d := time.Now().Add(shortDuration)
	//*创建有截止时间的context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	//*监听1s和截止时间context哪个先结束
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err(), time.Now()) //*1ms先结束
	}
}
