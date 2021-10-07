package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveComment(test *testing.T) {
	type args struct {
		input string
	}

	testsCases := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string with comment",
			args: args{input: "test1 // test2"},
			want: "test1 ",
		},
		{
			name: "string without comment",
			args: args{input: "test1"},
			want: "test1",
		},
		{
			name: "only comment",
			args: args{input: "// test2"},
			want: "",
		},
		{
			name: "empty string",
			args: args{input: ""},
			want: "",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			got := RemoveComment(testCase.args.input)

			assert.Equal(test, testCase.want, got)
		})
	}
}
