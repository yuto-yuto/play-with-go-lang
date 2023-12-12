package utils

import (
	"context"
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

type Debouncer2 struct {
	Timeout      time.Duration
	ForceTimeout time.Duration
	Callback     func()
	timer        *time.Timer
	startedTime  *time.Time
	mutex        sync.Mutex
}

func NewDebouncer2(timeout, forceTimeout time.Duration, callback func()) *Debouncer2 {
	return &Debouncer2{
		Timeout:      timeout,
		ForceTimeout: forceTimeout,
		Callback:     callback,
	}
}

func (m *Debouncer2) Debounce() {
	if m.timer == nil {
		m.assignTimerOnlyOnce()

		return
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.startedTime != nil && time.Since(*m.startedTime) > m.ForceTimeout {
		m.startedTime = nil
		m.Callback()

		return
	}

	now := time.Now()
	if !m.timer.Stop() {
		// It's unnecessary to drain timer.C because it's not used by time.AfterFunc
		m.startedTime = &now
	} else {
		if m.startedTime == nil {
			// Stop() returns true if callback is fired by force timeout too.
			m.startedTime = &now
		}
	}

	m.timer.Reset(m.Timeout)
}

func (m *Debouncer2) assignTimerOnlyOnce() {
	callback := func() {
		m.mutex.Lock()
		defer m.mutex.Unlock()

		if m.startedTime != nil {
			m.startedTime = nil
			m.Callback()
		}
	}

	m.timer = time.AfterFunc(m.Timeout, callback)
	now := time.Now()
	m.startedTime = &now
}

type Debouncer3 struct {
	timeChan chan time.Time
}

func NewDebouncer3(ctx context.Context, timeout, forceTimeout time.Duration, callback func()) *Debouncer3 {
	result := &Debouncer3{
		timeChan: make(chan time.Time, 100),
	}

	go func() {
		var startedTime *time.Time

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(timeout):
				if len(result.timeChan) == 0 && startedTime == nil {
					continue
				}

				startedTime = nil

				callback()
			case calledTime := <-result.timeChan:
				if startedTime == nil {
					startedTime = &calledTime
				} else if time.Since(*startedTime) > forceTimeout {
					startedTime = nil
					callback()
				}
			}
		}
	}()

	return result
}

func (m *Debouncer3) Debounce() {
	m.timeChan <- time.Now()
}

type Debouncer4 struct {
	timeChan chan time.Time
}

func NewDebouncer4(ctx context.Context, timeout, forceTimeout time.Duration, callback func()) *Debouncer4 {
	instance := &Debouncer4{
		timeChan: make(chan time.Time, 100),
	}

	go func() {
		var startedTime *time.Time
		var updatedTime *time.Time

		ticker := time.NewTicker(timeout)
		ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if updatedTime == nil {
					ticker.Stop()
					startedTime = nil
					updatedTime = nil
					callback()

					continue
				}

				now := time.Now()
				diffForce := startedTime.Add(forceTimeout).Sub(now)
				diffNormal := updatedTime.Add(timeout).Sub(now)

				diff := diffNormal
				if diffForce < diffNormal {
					diff = diffForce
				}

				if diff <= 0 {
					ticker.Stop()
					startedTime = nil
					updatedTime = nil
					callback()

					continue

				}

				ticker.Reset(diff)
			case timestamp := <-instance.timeChan:
				if startedTime == nil {
					startedTime = &timestamp
					ticker.Reset(timeout)
				} else {
					updatedTime = &timestamp
				}
			}
		}
	}()

	return instance

}

func (m *Debouncer4) Debounce() {
	m.timeChan <- time.Now()
}
