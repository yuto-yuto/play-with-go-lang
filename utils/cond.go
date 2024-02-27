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
	s.signal <- true
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
		withChannel.Update(0)
	}

	fmt.Println("--- Update state to 1 ---")
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("send channel %d\n", i)
		withChannel.Update(1)
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

func (c *condTester) SignalUpdate(value int) {
	c.update(value)
	c.cond.Signal()
}

func (c *condTester) BroadcastUpdate(value int) {
	c.update(value)
	c.cond.Broadcast()
}

func (c *condTester) update(value int) {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	c.state = value
}

func withCond() {
	condTester := &condTester{}
	condTester.cond = sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Waiting %d\n", i)
			condTester.Wait()
			fmt.Printf("Go %d\n", i)
			wg.Done()
		}(i)
	}

	fmt.Println("--- Waiting for a second ---")
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("send signal %d\n", i)
		condTester.SignalUpdate(0)
	}

	fmt.Println("--- Signal Update state to 1 ---")

	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("send signal %d\n", i)
		condTester.SignalUpdate(1)
	}
	
	
	fmt.Println("--- Signal with 0 ---")
	condTester.SignalUpdate(0)

	time.Sleep(2 * time.Second)
	fmt.Println("--- Broadcast ---")
	condTester.BroadcastUpdate(1)
	wg.Wait()
}

func badCondExample() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Waiting %d\n", i)
			mutex.Lock()
			defer mutex.Unlock()
			cond.Wait()

			fmt.Printf("Start %d\n", i)
			time.Sleep(time.Second)
			fmt.Printf("End %d\n", i)

			wg.Done()
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("--- Broadcast ---")
	cond.Broadcast()
	wg.Wait()
}

func RunCond() {
	withCond()
	// withChannel()
	// badCondExample()
}
