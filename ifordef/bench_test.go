package ifordef

import (
	"testing"
)

func BenchmarkDef(b *testing.B) {
	fn := NewDef(func() bool {
		return true
	})
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		if fn.Call() {
			count++
		}
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}

func BenchmarkIf(b *testing.B) {
	fn := NewIf(func() bool {
		return true
	})
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		if fn.Call() {
			count++
		}
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}
