package maporarray

import (
	"strconv"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	list := make(map[string]struct{}, b.N)
	for i := 0; i < b.N; i++ {
		idx := strconv.FormatInt(int64(i), 10)
		list[idx] = struct{}{}
	}

	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		idx := strconv.FormatInt(int64(i), 10)
		if _, ok := list[idx]; ok {
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
	list := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = strconv.FormatInt(int64(i), 10)
	}
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		idx := strconv.FormatInt(int64(i), 10)
		if list[i] == idx {
			count++
		}
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}
