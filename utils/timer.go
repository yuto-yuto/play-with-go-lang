package utils

import (
	"fmt"
	"sync"
	"time"
)

func RunTimer() {
	runTimer1()
	runTimer2()
	runTimer3()
	runDebounce()
}

func runTimer1() {
	timer := time.NewTimer(time.Second)

	var wg sync.WaitGroup
	wg.Add(1)

	go func(timer *time.Timer) {
		<-timer.C
		fmt.Println("this is executed after 1 second")
		wg.Done()
	}(timer)

	wg.Wait()
}

func runTimer2() {
	var wg sync.WaitGroup
	wg.Add(1)

	callback := func() {
		fmt.Println("this is executed after 1 second")
		wg.Done()
	}

	timer := time.AfterFunc(time.Second, callback)

	wg.Wait()
	fmt.Println(<-timer.C)
}

func runTimer3() {
	calledTime := time.Now()
	timeout := time.Millisecond * 500
	timer := time.NewTimer(time.Second)

	var wg sync.WaitGroup
	wg.Add(1)

	go func(timer *time.Timer) {
		<-timer.C
		fmt.Printf("called after %f second\n", time.Since(calledTime).Seconds())
		wg.Done()
	}(timer)

	time.Sleep(timeout)
	timer.Stop()
	timer.Reset(time.Second)

	time.Sleep(timeout)
	timer.Stop()
	timer.Reset(time.Second)

	wg.Wait()
}

func runDebounce() {
	var wg sync.WaitGroup
	wg.Add(1)

	debounceCallCount := 0
	callback := func() {
		fmt.Printf("Debounce call count: %d\n", debounceCallCount)
		wg.Done()
	}
	debouncer := NewDebounce(time.Second, callback)

	debounceCallCount++
	debouncer.Debounce()
	time.Sleep(500 * time.Millisecond)

	debounceCallCount++
	debouncer.Debounce()
	time.Sleep(500 * time.Millisecond)

	debounceCallCount++
	debouncer.Debounce()
	time.Sleep(500 * time.Millisecond)

	wg.Wait()
}

type Debouncer struct {
	timeout  time.Duration
	timer    *time.Timer
	callback func()
	mutex    sync.Mutex
}

func NewDebounce(timeout time.Duration, callback func()) Debouncer {
	return Debouncer{
		timeout:  timeout,
		callback: callback,
	}
}

func (m *Debouncer) Debounce() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.timer == nil {
		m.timer = time.AfterFunc(m.timeout, m.callback)

		return
	}

	m.timer.Stop()
	m.timer.Reset(m.timeout)
}

func (m *Debouncer) UpdateDebounceCallback(callback func()) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.timer.Stop()
	m.timer = time.AfterFunc(m.timeout, callback)
}
