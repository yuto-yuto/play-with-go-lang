package benchmark

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringConcatenation(b *testing.B) {
	var str string
	for i := 0; i < 10000; i++ {
		str = str + "abc"
	}
}

func BenchmarkStringConcatenation2(b *testing.B) {
	var str string
	for i := 0; i < 10000; i++ {
		str = fmt.Sprintf("%sabc", str)
	}
}

func BenchmarkStringConcatenation3(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < 10000; i++ {
		sb.WriteString("abc")
	}
}

func BenchmarkStringConcatenation4(b *testing.B) {
	var stringBytes []byte
	for i := 0; i < 10000; i++ {
		stringBytes = append(stringBytes, "abc"...)
	}
}

// go test ./benchmark -bench String -benchmem
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/benchmark
// cpu: Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz
// BenchmarkStringConcatenation-12         1000000000               0.02485 ns/op         0 B/op          0 allocs/op
// BenchmarkStringConcatenation2-12        1000000000               0.02831 ns/op         0 B/op          0 allocs/op
// BenchmarkStringConcatenation3-12        1000000000               0.0000256 ns/op               0 B/op          0 allocs/op
// BenchmarkStringConcatenation4-12        1000000000               0.0000412 ns/op               0 B/op          0 allocs/op
// PASS
// ok      play-with-go-lang/benchmark     0.406s