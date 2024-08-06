package ifordef

import (
	"sync"
)

type Mutex struct {
	sync.Mutex
	values []int
}

func NewMutex(size int) *Mutex {
	var values []int
	for i := 0; i < size; i++ {
		values = append(values, i)
	}
	return &Mutex{
		values: values,
	}
}

func (m *Mutex) Read(idx int) int {
	m.Lock()
	defer m.Unlock()
	return m.values[idx]
}
