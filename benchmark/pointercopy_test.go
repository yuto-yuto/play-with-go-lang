package utils

import (
	"testing"
)

type pair struct {
	a, b int
}

type bigData struct {
	prop1 string
	prop2 string
	prop3 int
	prop4 int
	prop5 int
	prop6 int
}

func NewPairValue() pair {
	return pair{1, 2}
}
func NewPairPointer() *pair {
	return &pair{1, 2}
}

func NewBigData() bigData {
	return bigData{
		prop1: "123456789012345678901234567890",
		prop2: "123456789012345678901234567890",
		prop3: 11,
		prop4: 22,
		prop5: 33,
		prop6: 44,
	}
}

func NewBigDataPointer() *bigData {
	return &bigData{
		prop1: "123456789012345678901234567890",
		prop2: "123456789012345678901234567890",
		prop3: 11,
		prop4: 22,
		prop5: 33,
		prop6: 44,
	}
}

func BenchmarkBigDataNewValue(b *testing.B) {
	var _ bigData
	for i := 0; i < b.N; i++ {
		_ = NewBigData()
	}
}

func BenchmarkBigDataValueCopy(b *testing.B) {
	var _ bigData
	var instance = NewBigData()
	for i := 0; i < b.N; i++ {
		_ = instance
	}
}

func BenchmarkBigDataNewPointer(b *testing.B) {
	var _ *bigData
	for i := 0; i < b.N; i++ {
		_ = NewBigDataPointer()
	}
}

func BenchmarkBigDataPointerCopy(b *testing.B) {
	var _ *bigData
	var instance = NewBigDataPointer()
	for i := 0; i < b.N; i++ {
		_ = instance
	}
}

func BenchmarkPairNewValue(b *testing.B) {
	var _ pair
	for i := 0; i < b.N; i++ {
		_ = NewPairValue()
	}
}

func BenchmarkPairValueCopy(b *testing.B) {
	var _ pair
	var instance = NewPairValue()
	for i := 0; i < b.N; i++ {
		_ = instance
	}
}

func BenchmarkPairNewPointer(b *testing.B) {
	var _ *pair
	for i := 0; i < b.N; i++ {
		_ = NewPairPointer()
	}
}

func BenchmarkPairPointerCopy(b *testing.B) {
	var _ *pair
	var instance = NewPairPointer()
	for i := 0; i < b.N; i++ {
		_ = instance
	}
}

func BenchmarkAppendWithoutSize(b *testing.B) {
	array := make([]int, 0)
	for i := 0; i < b.N; i++ {
		array = append(array, i)
	}
}
func BenchmarkAppendWithSize(b *testing.B) {
	array := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		array = append(array, i)
	}
}

// vscode ➜ /workspaces/play-with-go-lang/utils (main) $ go test -bench .
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/utils
// cpu: Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz
// BenchmarkBigDataNewValue-12             1000000000               0.2499 ns/op
// BenchmarkBigDataValueCopy-12            1000000000               0.2677 ns/op
// BenchmarkBigDataNewPointer-12           1000000000               0.2799 ns/op
// BenchmarkBigDataPointerCopy-12          1000000000               0.2536 ns/op
// BenchmarkPairNewValue-12                1000000000               0.2542 ns/op
// BenchmarkPairValueCopy-12               1000000000               0.2699 ns/op
// BenchmarkPairNewPointer-12              1000000000               0.2659 ns/op
// BenchmarkPairPointerCopy-12             1000000000               0.2944 ns/op
// BenchmarkAppendWithoutSize-12           203014581               18.64 ns/op
// BenchmarkAppendWithSize-12              1000000000               8.135 ns/op

// vscode ➜ /workspaces/play-with-go/utils (main) $ go test -bench .
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/utils
// cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
// BenchmarkBigDataNewValue-4              1000000000               0.3775 ns/op
// BenchmarkBigDataValueCopy-4             1000000000               0.3333 ns/op
// BenchmarkBigDataNewPointer-4            1000000000               0.3451 ns/op
// BenchmarkBigDataPointerCopy-4           1000000000               0.3327 ns/op
// BenchmarkPairNewValue-4                 1000000000               0.3370 ns/op
// BenchmarkPairValueCopy-4                1000000000               0.3433 ns/op
// BenchmarkPairNewPointer-4               1000000000               0.3312 ns/op
// BenchmarkPairPointerCopy-4              1000000000               0.4139 ns/op
// BenchmarkAppendWithoutSize-4            50692158               123.3 ns/op
// BenchmarkAppendWithSize-4       		500221972              111.4 ns/op
// PASS
// ok      play-with-go-lang/utils 65.758s

// vscode ➜ /workspaces/play-with-go-lang/utils (main) $ go test -bench . -benchmem
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/utils
// cpu: Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz
// BenchmarkBigDataNewValue-12             1000000000               0.2927 ns/op          0 B/op          0 allocs/op
// BenchmarkBigDataValueCopy-12            1000000000               0.2595 ns/op          0 B/op          0 allocs/op
// BenchmarkBigDataNewPointer-12           1000000000               0.2637 ns/op          0 B/op          0 allocs/op
// BenchmarkBigDataPointerCopy-12          1000000000               0.2784 ns/op          0 B/op          0 allocs/op
// BenchmarkPairNewValue-12                1000000000               0.3002 ns/op          0 B/op          0 allocs/op
// BenchmarkPairValueCopy-12               1000000000               0.3335 ns/op          0 B/op          0 allocs/op
// BenchmarkPairNewPointer-12              1000000000               0.2608 ns/op          0 B/op          0 allocs/op
// BenchmarkPairPointerCopy-12             1000000000               0.2851 ns/op          0 B/op          0 allocs/op
// BenchmarkAppendWithoutSize-12           100000000               11.06 ns/op           45 B/op          0 allocs/op
// BenchmarkAppendWithSize-12              1000000000               2.059 ns/op           8 B/op          0 allocs/op
// PASS
// ok      play-with-go-lang/utils 5.857s

// vscode ➜ /workspaces/play-with-go/utils (main ✗) $ go test -bench . -benchmem
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/utils
// cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
// BenchmarkBigDataNewValue-4              1000000000               0.3333 ns/op          0 B/op          0 allocs/op
// BenchmarkBigDataValueCopy-4             1000000000               0.3450 ns/op          0 B/op          0 allocs/op
// BenchmarkBigDataNewPointer-4            1000000000               0.3873 ns/op          0 B/op          0 allocs/op
// BenchmarkBigDataPointerCopy-4           1000000000               0.3355 ns/op          0 B/op          0 allocs/op
// BenchmarkPairNewValue-4                 1000000000               0.3323 ns/op          0 B/op          0 allocs/op
// BenchmarkPairValueCopy-4                1000000000               0.3319 ns/op          0 B/op          0 allocs/op
// BenchmarkPairNewPointer-4               1000000000               0.3353 ns/op          0 B/op          0 allocs/op
// BenchmarkPairPointerCopy-4              1000000000               0.3333 ns/op          0 B/op          0 allocs/op
// BenchmarkAppendWithoutSize-4            104694237               22.64 ns/op           43 B/op          0 allocs/op
// BenchmarkAppendWithSize-4               326078538               11.39 ns/op            8 B/op          0 allocs/op
// PASS
// ok      play-with-go-lang/utils 10.665s