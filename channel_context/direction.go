package channelcontext

import (
	"fmt"
	"time"
)

type channelBidirection struct {
	Bidirectional chan int
}

func newChannelBidirection() *channelBidirection {
	return &channelBidirection{
		Bidirectional: make(chan int, 3),
	}
}

func (c *channelBidirection) run(baseValue int) {
	limit := baseValue + 3

	for i := baseValue; i < limit; i++ {
		c.Bidirectional <- i
	}
}

type channelWithDirection struct {
	Receiver      <-chan int
	sender        chan<- int
}

func newChanelwithDirection() *channelWithDirection {
	c := make(chan int, 3)

	return &channelWithDirection{
		Receiver:      c,
		sender:        c,
	}
}

func (c *channelWithDirection) getReceive() <-chan int {
	return c.Receiver
}

func (c *channelWithDirection) tryToSend(value int) {
	select {
	case c.sender <- value:
	default:
		fmt.Printf("failed to send a data because buffer size was full. Data: [%d]\n", value)
	}
}

func (c *channelWithDirection) sendWithBlock(value int) {
	c.sender <- value
}

func runChannelBidirection() {
	middleMan := newChannelBidirection()

	fmt.Println("--- self send/receive---")
	middleMan.Bidirectional <- 11
	middleMan.Bidirectional <- 12
	middleMan.Bidirectional <- 13
	fmt.Println(<-middleMan.Bidirectional)
	fmt.Println(<-middleMan.Bidirectional)
	fmt.Println(<-middleMan.Bidirectional)

	fmt.Println("---TryToSend/receive---")
	go middleMan.run(9)
	fmt.Println(<-middleMan.Bidirectional)
	fmt.Println(<-middleMan.Bidirectional)
	fmt.Println(<-middleMan.Bidirectional)
}

func runChannelWithDirection() {
	sleepTime := 10 * time.Millisecond
	middleMan := newChanelwithDirection()

	// invalid operation: cannot send to receive-only channel middleMan.Receiver (variable of type <-chan int)
	// middleMan.Receiver <- 1

	go func() {
		loopCount := 5
		for i := 0; i < loopCount; i++ {
			middleMan.tryToSend(i)
		}

		for i := loopCount; i < 2*loopCount; i++ {
			middleMan.sendWithBlock(i)
		}
	}()

	for i := 0; i < 5; i++ {
		data := <-middleMan.Receiver
		fmt.Printf("Receiver(%d): %d\n", i, data)
		time.Sleep(sleepTime)
	}

	for i := 0; i < 5; i++ {
		select {
		case data := <-middleMan.getReceive():
			fmt.Printf("GetReceiver(%d): %d\n", i, data)
		default:
			fmt.Println("no more value")
		}
		time.Sleep(sleepTime)
	}
}

func RunChannelDirection() {
	fmt.Println("--- RunChannelBidirection ---")
	runChannelBidirection()

	fmt.Println("--- RunChannelWithDirection ---")
	runChannelWithDirection()
}
