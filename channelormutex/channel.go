package ifordef

import (
	"sync"
)

type Channel struct {
	values []int
	pool   sync.Pool
	chanR  chan *Result
}

type Result struct {
	Res chan int
	Idx int
}

func NewChannel(size int) *Channel {
	var values []int
	for i := 0; i < size; i++ {
		values = append(values, i)
	}
	C := &Channel{
		values: values,
		chanR:  make(chan *Result, 1),
		pool: sync.Pool{
			New: func() interface{} {
				return &Result{
					Res: make(chan int),
				}
			},
		},
	}
	go C.handle()
	return C
}

func (c *Channel) Read(idx int) int {
	r := c.pool.Get().(*Result)
	defer c.pool.Put(r)
	r.Idx = idx
	c.chanR <- r
	return <-r.Res
}

func (c *Channel) handle() {
	for {
		select {
		case r, ok := <-c.chanR:
			if !ok {
				return
			}
			r.Res <- c.values[r.Idx]
		}
	}
}
