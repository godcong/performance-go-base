package maporarray

import (
	"testing"
)

func BenchmarkMap(b *testing.B) {
	list := make(map[int]struct{}, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = struct{}{}
	}

	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		if _, ok := list[i]; ok {
			count++
		}
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}

func BenchmarkArrKey(b *testing.B) {
	list := make([]struct{}, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = struct{}{}
	}
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		if list[i] == struct{}{} {
			count++
		}
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}

func BenchmarkArrVal(b *testing.B) {
	list := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = i
	}
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		for idx := range list {
			if list[idx] == i {
				count++
				break
			}
		}
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}
