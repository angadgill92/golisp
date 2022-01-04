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

func Test_atom(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want Atom
	}{
		{
			name: "Parse integer token as a number",
			args: args{token: "1"},
			want: Number(1),
		},
		{
			name: "Parse open parens as a Symbol",
			args: args{token: "("},
			want: Symbol("("),
		},
		{
			name: "Parse closing parens as symbol",
			args: args{token: ")"},
			want: Symbol(")"),
		},
		{
			name: "Parse word token as a Symbol",
			args: args{token: "abc"},
			want: Symbol("abc"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := atom(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("atom() = %v, want %v", got, tt.want)
			}
		})
	}
}
