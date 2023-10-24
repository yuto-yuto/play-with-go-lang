package utils

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func RunAtomic() {
	var intValue atomic.Value
	fmt.Println(intValue.Load() == nil) // true

	fmt.Println("atomic: ", WriteInGoroutine())
	fmt.Println("mutex: ", WriteInGoroutineWithMutex())

	// storeDifferentType()
	atomicWriteOrder()
	atomicPointer()

}

func WriteInGoroutine() int64 {
	var ops atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return ops.Load()
}

func WriteInGoroutine2() int64 {
	var ops int64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddInt64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return ops
}

func WriteInGoroutineWithMutex() int64 {
	var ops int64
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				mutex.Lock()
				ops++
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return ops
}

func storeDifferentType() {
	fmt.Println("--- storeDifferentType ---")

	var intValue atomic.Value
	fmt.Println(intValue.Load())
	myData := myStruct{name: "Yuto", age: 36}

	intValue.Store(myData)
	loadedData := intValue.Load()
	fmt.Println(loadedData)
	fmt.Println(loadedData.(myStruct))

	person := person{
		Name:   "Yuto",
		Age:    36,
		Gender: "man",
	}
	// panic: sync/atomic: swap of inconsistently typed value into Value
	intValue.Store(person)
	fmt.Println(intValue.Load())
}

func atomicWriteOrder() {
	fmt.Println("--- atomicWriteOrder ---")
	var intValue atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		val := int64(i)
		go func() {
			intValue.Store(val)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(intValue.Load())
}

func atomicPointer() {
	fmt.Println("--- atomicPointer ---")

	var pointerValue atomic.Pointer[myStruct]
	fmt.Println(pointerValue.Load() == nil)

	old := myStruct{name: "name-1", age: 1}
	pointerValue.Store(&old)
	fmt.Println("whole struct: ", pointerValue.Load())
	fmt.Println("name: ", pointerValue.Load().name)
	fmt.Println("age: ", pointerValue.Load().age)

	fmt.Println("--- CompareAndSwap ---")
	newData := myStruct{name: "name-1", age: 2}
	pointerValue.CompareAndSwap(&old, &newData)
	fmt.Println("whole struct: ", pointerValue.Load())
	fmt.Println("name: ", pointerValue.Load().name)
	fmt.Println("age: ", pointerValue.Load().age)
}
