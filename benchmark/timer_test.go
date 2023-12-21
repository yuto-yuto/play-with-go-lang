package benchmark_test

import (
	"testing"
	"time"
)

func BenchmarkTimerAfterFunc(b *testing.B) {
	timer := time.AfterFunc(time.Hour, func() {})

	for i := 0; i < b.N; i++ {
		timer.Stop()
		timer = time.AfterFunc(time.Hour, func() {})
	}
}

func BenchmarkTimerTimer(b *testing.B) {
	timer := time.NewTimer(time.Hour)

	for i := 0; i < b.N; i++ {
		timer.Stop()
		timer.Reset(time.Hour)
	}
}

func BenchmarkTimerTimerWithIf(b *testing.B) {
	timer := time.NewTimer(time.Hour)

	for i := 0; i < b.N; i++ {
		if !timer.Stop() {
			<-timer.C
		}
		timer.Reset(time.Hour)
	}
}

func BenchmarkTimerTicker(b *testing.B) {
	ticker := time.NewTicker(time.Hour)
	ticker.Stop()

	for i := 0; i < b.N; i++ {
		ticker.Reset(time.Hour)
	}
}

func BenchmarkTimerTicker2(b *testing.B) {
	ticker := time.NewTicker(time.Hour)

	for i := 0; i < b.N; i++ {
		ticker.Stop()
		ticker.Reset(time.Hour)
	}
}

func BenchmarkTimerAfter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <-time.After(time.Nanosecond):
		case <-time.After(time.Millisecond):
		}
	}
}

// $ go test ./benchmark -benchmem -bench Timer
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/benchmark
// cpu: Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz
// BenchmarkTimerAfterFunc-12               8241452               150.3 ns/op            80 B/op          1 allocs/op
// BenchmarkTimerTimer-12                  18680098                62.07 ns/op            0 B/op          0 allocs/op
// BenchmarkTimerTimerWithIf-12            18710127                61.95 ns/op            0 B/op          0 allocs/op
// BenchmarkTimerTicker-12                 16981581                63.44 ns/op            0 B/op          0 allocs/op
// BenchmarkTimerAfter-12                   1000000              1104 ns/op             432 B/op          6 allocs/op
// PASS
// ok      play-with-go-lang/benchmark     6.126s
