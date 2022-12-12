package channelcontext

import (
	"context"
	"fmt"
	"time"
)

const numberOfMsg = 3

func send(channel chan int) {
	for i := 0; i < numberOfMsg; i++ {
		channel <- (i + 1)
		time.Sleep(time.Second)
	}
}

func receive(channel chan int) {
	for {
		data := <-channel
		fmt.Printf("Received: %d\n", data)
		if data == numberOfMsg {
			break
		}
	}
}

func sendWithContext(ctx context.Context, channel chan int) {
	for i := 0; ; i++ {
		channel <- (i + 1)
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			fmt.Println("send loop ends")
			return
		default:
		}
	}
}

func receiveWithContext(ctx context.Context, channel chan int) {
	for {
		data := <-channel
		fmt.Printf("Received: %d\n", data)

		select {
		case <-ctx.Done():
			fmt.Println("receive loop ends")
			return
		default:
		}
	}
}

func receiveWithTimeoutContext(ctx context.Context, channel chan int) {

	for {
		select {
		case data := <-channel:
			fmt.Printf("Received: %d\n", data)
		case <-time.After(time.Second):
			fmt.Println("timeout")
		case <-ctx.Done():
			fmt.Println("canceled")
			return
		}
	}
}
