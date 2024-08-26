package app

import (
	"sync/atomic"
	"unsafe"
)

const BUF_LEN = 1 << 27 // 128 Mib

type BufferPool struct {
	buffers [10][]byte
	num     []int64
}

func NewBufferPool() *BufferPool {
	return &BufferPool{
		buffers: [10][]byte{},
		num:     make([]int64, 0, 20),
	}
}

func (b *BufferPool) Get(returns int) []byte {
	for i := range b.num {
		if atomic.CompareAndSwapInt64(&b.num[i], 0, int64(returns)) {
			return b.buffers[i]
		}
	}
	if len(b.num) >= 10 {
		panic("too much buffers")
	}
	newBytes := make([]byte, BUF_LEN)

	b.num = append(b.num, int64(returns))
	b.buffers[len(b.num)-1] = newBytes
	return newBytes
}

func (b *BufferPool) Ret(bufPart []byte) {
	for i := range b.buffers {
		// check if  &b.bytes[i][0]<= &bufPart[0] < &b.bytes[i][0] + BUF_LEN
		d := uintptr(unsafe.Pointer(&bufPart[0])) - uintptr(unsafe.Pointer(&b.buffers[i][0]))
		if d >= 0 && d < BUF_LEN {
			atomic.AddInt64(&b.num[i], -1)
			return
		}
	}
	panic("buffer not found")
}
