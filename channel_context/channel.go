package channelcontext

import (
	"fmt"
	"time"
)

func runChannelTest() {
	channel := make(chan string)

	go send(channel)
	go receive(channel)

	fmt.Println("Do something...")
	time.Sleep((numberOfMsg + 1) * time.Second)
}

func runChannelTest2() {
	channel := make(chan string)

	send := func(channel chan string, data string, delay time.Duration) {
		time.Sleep(delay)
		channel <- data
	}
	go send(channel, "Hello 1", 100*time.Millisecond)
	go send(channel, "Hello 2", 2000*time.Millisecond)
	go send(channel, "Hello 3", 500*time.Millisecond)

	timeout := time.After(1000 * time.Millisecond)
	var results []string
	for i := 0; i < 3; i++ {
		select {
		case result := <-channel:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
		}
	}

	fmt.Println(results)
}

func runChannelTest3() {
	channel := make(chan string)

	send := func(channel chan string, data string, delay time.Duration) {
		time.Sleep(delay)
		channel <- data
	}
	go send(channel, "Hello 1", 100*time.Millisecond)
	go send(channel, "Hello 2", 2000*time.Millisecond)
	go send(channel, "Hello 3", 50*time.Millisecond)

	result := <-channel

	fmt.Println(result)
}

func runChannelTest4() {
	channel := make(chan string)

	send := func(channel chan string, data string, delay time.Duration) {
		time.Sleep(delay)
		channel <- data
	}
	go send(channel, "Hello 1", 100*time.Millisecond)
	go send(channel, "Hello 2", 2000*time.Millisecond)
	go send(channel, "Hello 3", 50*time.Millisecond)

	timeout := time.After(10 * time.Millisecond)
	select {
	case result := <-channel:
		fmt.Println(result)
	case <-timeout:
		fmt.Println("timed out")
	}
}
