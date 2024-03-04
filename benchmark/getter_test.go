package benchmark_test

import "testing"

type manyInt64 struct {
	prop0 int64
	prop1 int64
	prop2 int64
	prop3 int64
	prop4 int64
	prop5 int64
	prop6 int64
	prop7 int64
	prop8 int64
	prop9 int64
}

type dataHolder struct {
	prop0 manyInt64
	prop1 manyInt64
	prop2 manyInt64
	prop3 manyInt64
	prop4 manyInt64
	prop5 manyInt64
	prop6 manyInt64
	prop7 manyInt64
	prop8 manyInt64
	prop9 manyInt64
}

type smallData struct {
	data int32
}

func (m smallData) ValueGetter() int32 {
	return m.data
}

func (m *smallData) PointerGetter() int32 {
	return m.data
}

func (m dataHolder) ValueProp0() manyInt64 {
	return m.prop0
}

func (m *dataHolder) PointerProp0() manyInt64 {
	return m.prop0
}

func (m dataHolder) ValueProp9() manyInt64 {
	return m.prop9
}

func (m *dataHolder) PointerProp9() manyInt64 {
	return m.prop9
}

func BenchmarkSmallStructValueGetter(b *testing.B) {
	var data smallData
	for i := 0; i < b.N; i++ {
		_ = data.ValueGetter()
	}
}

func BenchmarkSmallStructPointerGetter(b *testing.B) {
	var data smallData
	for i := 0; i < b.N; i++ {
		_ = data.PointerGetter()
	}
}

func BenchmarkValueGetter0(b *testing.B) {
	var data dataHolder
	for i := 0; i < b.N; i++ {
		_ = data.ValueProp0()
	}
}

func BenchmarkPointerGetter0(b *testing.B) {
	var data dataHolder
	for i := 0; i < b.N; i++ {
		_ = data.PointerProp0()
	}
}

func BenchmarkValueGetter9(b *testing.B) {
	var data dataHolder
	for i := 0; i < b.N; i++ {
		_ = data.ValueProp9()
	}
}

func BenchmarkPointerGetter9(b *testing.B) {
	var data dataHolder
	for i := 0; i < b.N; i++ {
		_ = data.PointerProp9()
	}
}

// $ go test ./benchmark -bench Getter -benchmem
// goos: linux
// goarch: amd64
// pkg: play-with-go-lang/benchmark
// cpu: Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz
// BenchmarkSmallStructValueGetter-12              1000000000               0.2555 ns/op          0 B/op          0 allocs/op
// BenchmarkSmallStructPointerGetter-12            1000000000               0.2419 ns/op          0 B/op          0 allocs/op
// BenchmarkValueGetter0-12                        88152307                12.98 ns/op            0 B/op          0 allocs/op
// BenchmarkPointerGetter0-12                      1000000000               0.4835 ns/op          0 B/op          0 allocs/op
// BenchmarkValueGetter9-12                        89456086                12.48 ns/op            0 B/op          0 allocs/op
// BenchmarkPointerGetter9-12                      1000000000               0.2505 ns/op          0 B/op          0 allocs/op
// PASS
// ok      play-with-go-lang/benchmark     3.667s
