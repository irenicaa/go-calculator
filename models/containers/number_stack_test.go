package containers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberStack_Push(test *testing.T) {
	type args struct {
		number float64
	}

	testsCases := []struct {
		name      string
		stack     NumberStack
		args      args
		wantStack NumberStack
	}{
		{
			name:      "nonempty",
			stack:     []float64{1.0, 2.0},
			args:      args{number: 3.0},
			wantStack: NumberStack{1.0, 2.0, 3.0},
		},
		{
			name:      "empty",
			stack:     []float64{},
			args:      args{number: 3.0},
			wantStack: NumberStack{3.0},
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			testCase.stack.Push(testCase.args.number)

			assert.Equal(test, testCase.wantStack, testCase.stack)
		})
	}
}

func TestNumberStack_Pop(test *testing.T) {
	testsCases := []struct {
		name       string
		stack      NumberStack
		wantNumber float64
		wantOk     bool
	}{
		{
			name:       "nonempty",
			stack:      []float64{1.0, 2.0},
			wantNumber: 2.0,
			wantOk:     true,
		},
		{
			name:       "empty",
			stack:      []float64{},
			wantNumber: 0,
			wantOk:     false,
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotToken, gotOk := testCase.stack.Pop()

			assert.Equal(test, testCase.wantNumber, gotToken)
			assert.Equal(test, testCase.wantOk, gotOk)
		})
	}
}
