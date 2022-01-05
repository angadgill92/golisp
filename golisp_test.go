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

func Test_readFromTokens(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name       string
		args       args
		want       []string
		want1      Exp
		wantErr    bool
		wantErrVal error
	}{
		{
			name:       "Return 'unexpected EOF' on empty input",
			args:       args{tokens: []string{}},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrVal: errUnexpectedEOF,
		},
		{
			name:       "Return 'unexpected )' if tokens start with closing parens",
			args:       args{tokens: []string{")"}},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrVal: errUnexpectedCloseParens,
		},
		{
			name:       "Return list of expressions for a simple list",
			args:       args{tokens: []string{"(", "define", "10", ")"}},
			want:       []string{},
			want1:      List{Symbol("define"), Number(10)},
			wantErr:    false,
			wantErrVal: nil,
		},
		{
			name:       "Return list of expressions for nested lists",
			args:       args{tokens: []string{"(", "begin", "(", "define", "r", "10", ")", "(", "*", "pi", "(", "*", "r", "r", ")", ")", ")"}},
			want:       []string{},
			want1:      List{Symbol("begin"), List{Symbol("define"), Symbol("r"), Number(10)}, List{Symbol("*"), Symbol("pi"), List{Symbol("*"), Symbol("r"), Symbol("r")}}},
			wantErr:    false,
			wantErrVal: nil,
		},
		{
			name:  "Return empty expression list on empty input",
			args:  args{tokens: []string{"(", ")"}},
			want:  []string{},
			want1: List(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := readFromTokens(tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFromTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && (err != tt.wantErrVal) {
				t.Errorf("readFromTokens() error got = '%v', want '%v'", err, tt.wantErrVal)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFromTokens() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("readFromTokens() got1 = %v %T, want %v %T", got1, got1, tt.want1, tt.want1)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		program string
	}
	tests := []struct {
		name    string
		args    args
		wantExp Exp
		wantErr bool
	}{
		{
			name:    "Return list of expressions for nested lists",
			args:    args{program: "(begin (define r 10) (* pi (* r r)))"},
			wantExp: List{Symbol("begin"), List{Symbol("define"), Symbol("r"), Number(10)}, List{Symbol("*"), Symbol("pi"), List{Symbol("*"), Symbol("r"), Symbol("r")}}},
			wantErr: false,
		},
		{
			name:    "Return error on unexpected input",
			args:    args{program: "(begin (define r 10) (* pi (* r r))))"},
			wantExp: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExp, err := parse(tt.args.program)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotExp, tt.wantExp) {
				t.Errorf("parse() = %v, want %v", gotExp, tt.wantExp)
			}
		})
	}
}
