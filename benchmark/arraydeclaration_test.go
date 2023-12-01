package benchmark_test

import (
	"testing"
)

func BenchmarkArrayEmptyDeclaration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []int{}
	}
}
func BenchmarkArrayEmptyDeclarationAndAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array := []int{}
		for j := 0; j < 100; j++ {
			array = append(array, j)
		}
	}
}
func BenchmarkArrayVarDeclaration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ []int
	}
}
func BenchmarkArrayVarDeclarationAndAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var array []int
		for j := 0; j < 100; j++ {
			array = append(array, j)
		}
	}
}

func BenchmarkArrayMakeDeclaration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]int, 0)
	}
}
func BenchmarkArrayMakeDeclarationAndAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array := make([]int, 0)
		for j := 0; j < 100; j++ {
			array = append(array, j)
		}
	}
}

func BenchmarkArraySizeDeclaration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = [100]int{}
	}
}
func BenchmarkArraySizeDeclarationAndAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array := [100]int{}
		for j := 0; j < 100; j++ {
			array[j] = j
		}
	}
}

func BenchmarkArrayMakeSizeDeclaration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]int, 0, 100)
	}
}
func BenchmarkArrayMakeSizeDeclarationAndAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array := make([]int, 0, 100)
		for j := 0; j < 100; j++ {
			array = append(array, j)
		}
	}
}
func BenchmarkArrayMakeSizeDeclarationAndAssign2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array := make([]int, 100, 100)
		for j := 0; j < 100; j++ {
			array[j] = j
		}
	}
}

// $ go test ./benchmark -benchmem -bench Array
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/benchmark
// cpu: Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz
// BenchmarkArrayEmptyDeclaration-12                       1000000000               0.2452 ns/op          0 B/op          0 allocs/op
// BenchmarkArrayEmptyDeclarationAndAssign-12               2182192               546.4 ns/op          2040 B/op          8 allocs/op
// BenchmarkArrayVarDeclaration-12                         1000000000               0.2551 ns/op          0 B/op          0 allocs/op
// BenchmarkArrayVarDeclarationAndAssign-12                 2128820               540.4 ns/op          2040 B/op          8 allocs/op
// BenchmarkArrayMakeDeclaration-12                        1000000000               0.2537 ns/op          0 B/op          0 allocs/op
// BenchmarkArrayMakeDeclarationAndAssign-12                2127392               551.1 ns/op          2040 B/op          8 allocs/op
// BenchmarkArraySizeDeclaration-12                        1000000000               0.2453 ns/op          0 B/op          0 allocs/op
// BenchmarkArraySizeDeclarationAndAssign-12               38434558                30.57 ns/op            0 B/op          0 allocs/op
// BenchmarkArrayMakeSizeDeclaration-12                    1000000000               0.2438 ns/op          0 B/op          0 allocs/op
// BenchmarkArrayMakeSizeDeclarationAndAssign-12           20733373                57.97 ns/op            0 B/op          0 allocs/op
// BenchmarkArrayMakeSizeDeclarationAndAssign2-12          38073238                30.45 ns/op            0 B/op          0 allocs/op
// PASS
// ok      play-with-go-lang/benchmark     10.282s