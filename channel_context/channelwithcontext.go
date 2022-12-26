package channelcontext

import (
	"context"
	"time"
)

func runChannelWithContext() {
	channel := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	go sendWithContext(ctx, channel)
	go receiveWithContext(ctx, channel)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

func runChannelWithTimeout() {
	channel := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	go sendWithContext(ctx, channel)
	// go receiveWithContext(ctx, channel)
	go receiveWithTimeout(ctx, channel)

	time.Sleep(5 * time.Second)
	cancel()
}

func runChannelWithCancelAndTimeout() {
	channel := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	go sendWithContextWithRandomTime(ctx, channel)
	go receiveWithCancelAndTimeout(ctx, channel)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
