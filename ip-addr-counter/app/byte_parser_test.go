package app_test

import (
	"fmt"
	"ipac/app"
	"math"
	"math/rand"

	"strconv"
	"testing"
)

func TestParseByte(t *testing.T) {
	for i := byte(0); i < math.MaxUint8; i++ {
		bytes := strconv.FormatInt(int64(i), 10)
		result := app.ParseByte([]byte(bytes))
		if result != i {
			t.Errorf("expected %d: result: %d\n", i, result)
		}
	}
}

func BenchmarkParseByte(b *testing.B) {
	b.StopTimer()
	fmt.Println(b.N)
	randBytes := make([][]byte, b.N)
	for i := range randBytes {
		randBytes[i] = []byte(strconv.FormatUint(rand.Uint64()%math.MaxUint8, 10))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		app.ParseByte(randBytes[i])
	}
}

func BenchmarkParseUint(b *testing.B) {
	b.StopTimer()
	fmt.Println(b.N)
	randBytes := make([][]byte, b.N)
	for i := range randBytes {
		randBytes[i] = []byte(strconv.FormatUint(rand.Uint64()%math.MaxUint8, 10))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		strconv.ParseUint(string(randBytes[i]), 10, 8)
	}
}
