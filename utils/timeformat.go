package utils

import (
	"fmt"
	"time"
)

func FormatTime() {
	now := time.Now()
	fmt.Println(now)
	
	fmt.Println(now.Format("20060102_030405"))
	fmt.Println(now.Format("yyyy"))
	fmt.Println(now.Format("20230123_120055"))
	fmt.Println(now.Format("2023012312:00:55.000"))
	fmt.Println(now.Format("2023012312:00:55.000Z"))
}
