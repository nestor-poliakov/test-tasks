package app_test

import (
	"ipac/app"
	"testing"
)

func TestBufferPool(t *testing.T) {
	p := app.NewBufferPool()
	b1 := p.Get(1)
	b2 := p.Get(2)
	if &b1[0] == &b2[0] {
		t.Errorf("got the same slice twice")
	}
	p.Ret(b1[len(b1)-1:])
	p.Ret(b2)
	p.Ret(b2)
	b3 := p.Get(1)
	if &b1[0] != &b3[0] {
		t.Errorf("slice not returned to pool")
	}
}
