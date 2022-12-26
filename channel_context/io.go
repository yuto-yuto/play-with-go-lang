package channelcontext

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
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

func sendWithContext(ctx context.Context, channel chan string) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("send loop ends %s\n", ctx.Err().Error())
			return
		default:
			channel <- fmt.Sprintf("Hello: %d", (i + 1))
			time.Sleep(time.Second)
		}
	}
}

func sendWithContextWithRandomTime(ctx context.Context, channel chan string) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("send loop ends %s\n", ctx.Err().Error())
			return
		default:
			channel <- fmt.Sprintf("Hello: %d", (i + 1))
			randomDelay := time.Duration(rand.Intn(2000)) * time.Millisecond
			time.Sleep(randomDelay)
		}
	}
}

func receiveWithContext(ctx context.Context, channel chan string) {
	for {
		select {
		case data := <-channel:
			fmt.Printf("Received: %s\n", data)
		case <-ctx.Done():
			fmt.Printf("receive loop ends: %s\n", ctx.Err().Error())
			return
		}
	}
}

func receiveWithTimeout(ctx context.Context, channel chan string) {
	for {
		select {
		case data := <-channel:
			fmt.Printf("Received: %s\n", data)
		case <-ctx.Done():
			err := ctx.Err()
			if errors.Is(err, context.Canceled) {
				fmt.Println("Canceled")
			} else if errors.Is(err, context.DeadlineExceeded) {
				fmt.Println("Timeout")
			}
			return
		}
	}
}

func receiveWithCancelAndTimeout(ctx context.Context, channel chan string) {
	for {
		select {
		case data := <-channel:
			fmt.Printf("Received: %s\n", data)
		case <-time.After(time.Second):
			fmt.Println("timeout")
		case <-ctx.Done():
			fmt.Printf("receive loop ends: %s\n", ctx.Err().Error())
			return
		}
	}
}
