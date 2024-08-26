package app_test

import (
	"bytes"
	"ipac/app"

	"testing"
)

func TestRunreader(t *testing.T) {
	b := make([]byte, 100)
	for i := 0; i < 100; i++ {
		b[i] = byte(i)
	}
	rr := app.NewRunreader(bytes.NewReader(b))
	toRead := make([]byte, 21)

	n, err := rr.Read(toRead)
	if err != nil {
		t.Errorf("Read error: %s", err)
	}
	if n != 21 {
		t.Errorf("Red: number of bytes read expected %d, got %d", 21, n)
	}
	if !bytes.Equal(b[:21], toRead) {
		t.Errorf("Read: expected %v, got %v", b[:21], toRead)
	}
	rr.Unread(1)
	n, err = rr.Read(toRead)
	if err != nil {
		t.Errorf("Read error: %s", err)
	}
	if n != 21 {
		t.Errorf("Red: number of bytes read expected %d, got %d", 21, n)
	}
	if !bytes.Equal(b[20:41], toRead) {
		t.Errorf("Read after Unread: expected %v, got %v", b[20:41], toRead)
	}
	toRead2 := make([]byte, 150)
	n, err = rr.Read(toRead2)
	if err != nil {
		t.Errorf("Read error: %s", err)
	}
	if n != 59 {
		t.Errorf("Red: number of bytes read expected %d, got %d", 59, n)
	}
	if !bytes.Equal(toRead2[:n], b[41:]) {
		t.Errorf("Read after Unread: expected %v, got %v", b[41:], toRead2)
	}

}
