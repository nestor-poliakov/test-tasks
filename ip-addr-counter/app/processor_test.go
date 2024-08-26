package app_test

import (
	"ipac/app"
	"strconv"
	"strings"
	"testing"
)

func Test_Processor_Result(t *testing.T) {
	data := []uint32{2, 5, 0}
	expectedResult := 3
	result := app.Result(data)
	if result != expectedResult {
		t.Errorf("%v expected: %d bits, got %d\n", data, expectedResult, result)
	}
}

func Test_Processor(t *testing.T) {
	tests := []struct {
		data    string
		threads int
		result  int
	}{
		{
			data:    "1.2.3.4\n2.222.2.2\n3.3.3.123\n1.2.3.4\n",
			threads: 8,
			result:  3,
		},
		{
			data:    "1.2.3.4\n2.222.2.2\n3.3.3.123\n1.2.3.4",
			threads: 8,
			result:  3,
		},
		{
			data:    "1.2.3.4\n2.222.2.2\n3.3.3.123\n1.2.3.4\n",
			threads: 1,
			result:  3,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			p := app.NewProcessor()
			p.Process(strings.NewReader(tt.data), tt.threads)
			result := p.Result()
			if result != tt.result {
				t.Errorf("expected %d, got %d", tt.result, result)
			}
		})
	}
}
