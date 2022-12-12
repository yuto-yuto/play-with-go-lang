package channelcontext

import (
	"time"
)

func runChannelTest() {
	channel := make(chan int)

	go send(channel)
	go receive(channel)

	time.Sleep((numberOfMsg + 1) * time.Second)
}
