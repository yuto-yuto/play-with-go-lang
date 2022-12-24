package channelcontext

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const numberOfMsg = 3

func sendWithCallback(cb func(data string)) {
	for i := 0; i < 3; i++ {
		data := fmt.Sprintf("Hello: %d", i+1)
		cb(data)
		time.Sleep(time.Second)
	}
}

func sendWithCallback2(cb func(data string), wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		data := fmt.Sprintf("Hello: %d", i+1)
		cb(data)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func receiver(data string) {
	fmt.Printf("Received: [%s]\n", data)
}

func send(channel chan string) {
	for i := 0; i < numberOfMsg; i++ {
		channel <- fmt.Sprintf("Hello: %d", (i + 1))
		time.Sleep(time.Second)
	}
}

func receive(channel chan string) {
	for {
		data := <-channel
		fmt.Printf("Received: [%s]\n", data)
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
