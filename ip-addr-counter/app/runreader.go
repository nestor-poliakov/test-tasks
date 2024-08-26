package app

import (
	"fmt"
	"io"
)

const RBUF_LEN = 20

type Runreader struct {
	r   io.Reader
	buf []byte
	n   int
}

func NewRunreader(r io.Reader) *Runreader {
	return &Runreader{
		r:   r,
		buf: make([]byte, RBUF_LEN),
	}
}

func (r *Runreader) Read(b []byte) (int, error) {
	copy(b, r.buf[len(r.buf)-r.n:])
	n, err := r.r.Read(b[r.n:])
	if n < 20 {
		copy(r.buf, r.buf[n:])
		copy(r.buf[RBUF_LEN-n:], b[r.n:])
	} else {
		copy(r.buf, b[r.n+n-20:])
	}
	readed := n + r.n
	r.n = 0
	return readed, err
}

func (r *Runreader) Unread(n int) {
	if n > RBUF_LEN-r.n {
		panic(fmt.Errorf("too much for unread"))
	}
	r.n += n
}
