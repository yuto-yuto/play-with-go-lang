package channelcontext

import (
	"fmt"
	"time"
)

type channelWithDirection struct {
	Receiver <-chan int
	sender   chan<- int
}

func newChanelwithDirection() *channelWithDirection {
	c := make(chan int, 3)

	return &channelWithDirection{
		Receiver: c,
		sender:   c,
	}
}

func (c *channelWithDirection) GetReceive() <-chan int {
	return c.Receiver
}

func (c *channelWithDirection) TryToSend(value int) {
	select {
	case c.sender <- value:
	default:
		fmt.Printf("failed to send a data because buffer size was full. Data: [%d]\n", value)
	}
}

func (c *channelWithDirection) SendWithBlock(value int) {
	c.sender <- value
}

func RunChannelWithDirection() {
	sleepTime := 10 * time.Millisecond
	middleMan := newChanelwithDirection()

	// invalid operation: cannot send to receive-only channel dataReceiver.Receiver (variable of type <-chan int)
	// dataReceiver.Receiver <- 1

	go func() {
		loopCount := 5
		for i := 0; i < loopCount; i++ {
			middleMan.TryToSend(i)
		}

		for i := loopCount; i < 2*loopCount; i++ {
			middleMan.SendWithBlock(i)
		}
	}()

	for i := 0; i < 5; i++ {
		data := <-middleMan.Receiver
		fmt.Printf("Receiver(%d): %d\n", i, data)
		time.Sleep(sleepTime)
	}

	for i := 0; i < 5; i++ {
		select {
		case data := <-middleMan.GetReceive():
			fmt.Printf("GetReceiver(%d): %d\n", i, data)
		default:
			fmt.Println("no more value")
		}
		time.Sleep(sleepTime)
	}
}
