package utils

import (
	"fmt"
	"sync"
	"time"
)

func RunOnce() {

	fmt.Println("--- call 10 times ---")
	var once sync.Once

	for i := range 10 {
		once.Do(func() {
			fmt.Println(i)
		})
	}

	fmt.Println("--- 10 goroutines---")
	var once2 sync.Once
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			startTime := time.Now()
			fmt.Printf("start(%d)\n", i)
			once2.Do(func() {
				fmt.Printf("%s : %d", time.Now().String(), i)
				time.Sleep(time.Second)
			})

			fmt.Printf("end(%d): elapsed time %d ms\n", i, time.Since(startTime).Milliseconds())
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("--- completed ---")
}
