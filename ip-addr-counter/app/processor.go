package app

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
)

type Processor struct {
	data []uint32
}

func NewProcessor() *Processor {
	return &Processor{
		data: make([]uint32, 1<<27), // 512 MiB
	}
}

func (p *Processor) Process(reader io.Reader, threads int) {
	bufPool := NewBufferPool()
	r := NewRunreader(reader)
	ch := make(chan []byte, threads)
	wg := sync.WaitGroup{}
	wg.Add(threads)
	for range threads {
		go p.process(ch, &wg, bufPool)
	}
	for {
		buf := bufPool.Get(threads)
		n, err := r.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(fmt.Errorf("process eror: read from reader: %w", err))
		}
		buf = buf[:n]
		if n == BUF_LEN {
			end := 0
			for i := len(buf) - 20; i < len(buf); i++ {
				if buf[i] == '\n' {
					r.Unread(len(buf) - i - 1)
					end = i
					break
				}
			}
			buf = buf[:end]
		} else if buf[len(buf)-1] != '\n' {
			buf = append(buf, '\n')
		}

		partLen := BUF_LEN / threads
		if len(buf) <= partLen {
			ch <- buf
			continue
		}

		prevEnd := 0
		for i := 1; i < threads; i++ {
			if i*partLen > len(buf) {
				break
			}
			for k := i * partLen; ; k++ {
				if buf[k] == '\n' {
					ch <- buf[prevEnd : k+1]
					prevEnd = k + 1
					break
				}
			}
		}
		ch <- buf[prevEnd:]
	}
	close(ch)
	wg.Wait()
	return
}

func (p *Processor) process(ch chan []byte, wg *sync.WaitGroup, pool *BufferPool) {
	for s := range ch {
		p.processIps(s)
		pool.Ret(s)
	}
	wg.Done()
}

func (p *Processor) processIps(ips []byte) {
	// 4 bytes of one ip addr
	var ipBytes [4]uint32
	prevDot := -1
	byteNum := 0
	for i := 0; i < len(ips); i++ {
		c := ips[i]
		if c >= '0' {
			continue
		}
		if c == '.' {
			ipBytes[byteNum] = uint32(ParseByte(ips[prevDot+1 : i]))
			prevDot = i
			byteNum++
			continue
		}
		if c == '\n' {
			n := ipBytes[0]<<24 | ipBytes[1]<<16 | ipBytes[2]<<8 | uint32(ParseByte(ips[prevDot+1:i]))
			// first 27 bits is index in p.data
			// last 5 bits is number of bit
			atomic.OrUint32(&p.data[n>>5], (1 << (n & 0b11111)))
			byteNum = 0
			prevDot = i
		}
	}
}

func (p *Processor) Result() (res int) {
	for _, d := range p.data {
		// 1s faster than for loop
		res += int(
			d&1 +
				(d>>1)&1 +
				(d>>2)&1 +
				(d>>3)&1 +
				(d>>4)&1 +
				(d>>5)&1 +
				(d>>6)&1 +
				(d>>7)&1 +
				(d>>8)&1 +
				(d>>9)&1 +
				(d>>10)&1 +
				(d>>11)&1 +
				(d>>12)&1 +
				(d>>13)&1 +
				(d>>14)&1 +
				(d>>15)&1 +
				(d>>16)&1 +
				(d>>17)&1 +
				(d>>18)&1 +
				(d>>19)&1 +
				(d>>20)&1 +
				(d>>21)&1 +
				(d>>22)&1 +
				(d>>23)&1 +
				(d>>24)&1 +
				(d>>25)&1 +
				(d>>26)&1 +
				(d>>27)&1 +
				(d>>28)&1 +
				(d>>29)&1 +
				(d>>30)&1 +
				(d>>31)&1 +
				(d>>32)&1)
	}
	return res
}
