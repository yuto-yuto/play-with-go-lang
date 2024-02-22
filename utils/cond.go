package utils

import (
	"fmt"
	"sync"
	"time"
)

type syncWithChannel struct {
	signal chan bool
	mutex  sync.Mutex
	state  int
}

func (s *syncWithChannel) Wait() {
	s.mutex.Lock()
	for s.state == 0 {
		s.mutex.Unlock()

		<-s.signal

		s.mutex.Lock()
	}
	s.mutex.Unlock()
}

func (s *syncWithChannel) Update(value int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.state = value
}

func (s *syncWithChannel) GetCurrentState() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.state
}

func withChannel() {
	withChannel := &syncWithChannel{
		signal: make(chan bool, 10),
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Waiting %d\n", i)
			withChannel.Wait()
			fmt.Printf("Go %d\n", i)
		}(i)
	}

	fmt.Println("--- Waiting for a second ---")
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("send channel %d\n", i)
		withChannel.signal <- true
	}

	fmt.Println("--- Update state to 1 ---")

	withChannel.Update(1)
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("send channel %d\n", i)
		withChannel.signal <- true
	}

	time.Sleep(2 * time.Second)
	fmt.Println("--- Close ---")
	close(withChannel.signal)
	time.Sleep(time.Second)
}

type condTester struct {
	cond  *sync.Cond
	state int
}

func (c *condTester) Wait() {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	for c.state == 0 {
		c.cond.Wait()
	}
}

func (c *condTester) Update(value int) {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	c.state = value
}

func (c *condTester) GetCurrentState() int {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	return c.state
}

func withCond() {
	condTester := &condTester{}
	condTester.cond = sync.NewCond(&sync.Mutex{})

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Waiting %d\n", i)
			condTester.Wait()
			fmt.Printf("Go %d\n", i)
		}(i)
	}

	fmt.Println("--- Waiting for a second ---")
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("send signal %d\n", i)
		condTester.cond.Signal()
	}

	fmt.Println("--- Update state to 1 ---")

	condTester.Update(1)
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("send signal %d\n", i)
		condTester.cond.Signal()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("--- Broadcast ---")
	condTester.cond.Broadcast()
	time.Sleep(time.Second)
}

func RunCond() {
	// withCond()
	withChannel()
}
