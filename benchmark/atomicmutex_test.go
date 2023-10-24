package utils

import (
	"play-with-go-lang/utils"
	"testing"
)

func BenchmarkWriteWithAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = utils.WriteInGoroutine()
	}
}

func BenchmarkWriteWithAtomic2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = utils.WriteInGoroutine2()
	}
}
func BenchmarkWriteWithMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = utils.WriteInGoroutineWithMutex()
	}
}

// $ go test ./benchmark -bench Write
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/benchmark
// cpu: Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz
// BenchmarkWriteWithAtomic-12         1464            817827 ns/op
// BenchmarkWriteWithMutex-12           324           3864598 ns/op
// PASS
// ok      play-with-go-lang/benchmark     2.911s
