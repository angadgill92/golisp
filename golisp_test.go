package main

import (
	"reflect"
	"testing"
)

func Test_tokenize(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Surround open parens with space",
			args: args{
				input: "(",
			},
			want: []string{
				"(",
			},
		},

		{
			name: "Surround close parens with space",
			args: args{
				input: ")",
			},
			want: []string{
				")",
			},
		},

		{
			name: "Split input on space and return a string slice of tokens",
			args: args{
				input: "(add 2 3)",
			},
			want: []string{
				"(", "add", "2", "3", ")",
			},
		},

		{
			name: "Ignore new line etc",
			args: args{
				input: `(add 
					          3 (add 2 3)
							  )`,
			},
			want: []string{
				"(", "add", "3", "(", "add", "2", "3", ")", ")",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tokenize(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
