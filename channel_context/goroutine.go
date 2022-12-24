package channelcontext

import (
	"fmt"
	"sync"
)

func runWithoutGoroutine() {
	sendWithCallback(receiver)
	fmt.Println("Do something...")
}

func runWithGoroutine() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sendWithCallback2(receiver, &wg)
	fmt.Println("Do something...")
	wg.Wait()
}
