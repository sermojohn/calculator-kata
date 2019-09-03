package ops

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

// no shared state between tests

func TestAdd(t *testing.T) {
	// parameterised tests, e.g. JUnitParams
	// table driven tests in Go
	tests := []struct {
		title  string
		param1 int64
		param2 int64
		expRes int64
	}{
		{
			"zero plus zero", 0, 0, 0,
		},
		{
			"one plus one", 1, 1, 2,
		},
		{
			"one plus two", 1, 2, 3,
		},
		{
			"two plus three", 2, 3, 5,
		},
		{
			"-2 plus -3", -2, -3, -5,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			// arrange step
			logger := &MockLogger{}
			c := NewCalc(logger)

			// act step
			res := c.Add(test.param1, test.param2)

			// assert step
			assert.Equal(t, res, test.expRes)
			assert.Assert(t, is.Equal(logger.execCounter, 1))
		})
	}

}

type MockLogger struct {
	execCounter int
}

func (ml *MockLogger) Log(args ...interface{}) {
	fmt.Println("mocLogger log called")
	ml.execCounter++
}

func BenchmarkAdd(b *testing.B) {
	c := &Calc{}

	for i := 0; i < b.N; i++ {
		c.Add(1, 2)
	}
}
