package number

import "testing"

func BenchmarkSumOfNumber(b *testing.B) {
	data := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sumOfNumber(data)
	}
	b.StopTimer()
}

func BenchmarkSumOfNumberRef(b *testing.B) {
	data := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sumOfNumberRef(data)
	}
	b.StopTimer()
}

func BenchmarkSumOfNumberRefMore(b *testing.B) {
	data := 1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sumOfNumberMore(data)
	}
	b.StopTimer()
}
