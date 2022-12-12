package channelcontext

import (
	"context"
	"time"
)

func runChannelWithContext() {
	channel := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go sendWithContext(ctx, channel)
	go receiveWithContext(ctx, channel)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

func runChannelWithTimeoutContext() {
	channel := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go receiveWithTimeoutContext(ctx, channel)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
