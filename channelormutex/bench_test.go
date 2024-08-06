package ifordef

import (
	"testing"
)

func BenchmarkNewChannel(b *testing.B) {
	c := NewChannel(b.N)
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		if c.Read(i) == i {
			count++
		}
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}

func BenchmarkMutex(b *testing.B) {
	m := NewMutex(b.N)
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		if m.Read(i) == i {
			count++
		}
	}
	if count != b.N {
		b.Fatal("count != b.N")
	}
}
