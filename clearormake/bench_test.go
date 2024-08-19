package ifordef

import (
	"testing"
)

func BenchmarkClear(b *testing.B) {
	v := make(map[int]int, b.N)
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		clear(v)
		for j := 0; j < i; j++ {
			v[j] = j
		}
		count++
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}

func BenchmarkMake(b *testing.B) {
	v := make(map[int]int, b.N)
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		v = make(map[int]int)
		for j := 0; j < i; j++ {
			v[j] = j
		}
		count++
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}
