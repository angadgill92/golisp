package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Exp represents a scheme expression which can be an Ltom
// or a list
type Exp interface {
	expression()
}

// Atom is a symbol or number
type Atom interface {
	Exp
	atom()
}

// Symbol is a type alias representing symbols in lisp.
// Anything that is not a number is a symbol
type Symbol string

func (s Symbol) atom()       {}
func (s Symbol) expression() {}

type Number float64

func (n Number) atom()       {}
func (n Number) expression() {}

type List []Exp

func (l List) expression() {}

var (
	errUnexpectedEOF         = fmt.Errorf("unexpected EOF")
	errUnexpectedCloseParens = fmt.Errorf("unexpected )")
)

const (
	openParens  = "("
	closeParens = ")"
)

func parse(program string) (exp Exp, err error) {
	tokens, exp, err := readFromTokens(tokenize(program))

	// there should be no tokens left if the program is
	// syntactically correct
	if len(tokens) != 0 {
		return nil, fmt.Errorf("unexpected tokens:%v expected EOF", tokens)
	}
	return
}

// readFromTokens reads an expression from a sequence of
// tokens
func readFromTokens(tokens []string) ([]string, Exp, error) {
	if len(tokens) == 0 {
		return nil, nil, errUnexpectedEOF
	}

	token, tokens := tokens[0], tokens[1:]

	switch token {
	case openParens:
		var (
			list List
			exp  Exp
			err  error
		)

		for len(tokens) == 0 || tokens[0] != closeParens {
			tokens, exp, err = readFromTokens(tokens)
			if err != nil {
				return nil, nil, err
			}
			list = append(list, exp)
		}
		tokens = tokens[1:]
		return tokens, list, nil
	case closeParens:
		return nil, nil, errUnexpectedCloseParens
	}

	return tokens, atom(token), nil
}

// atom tries to parse a token as a Number (type alias for float64)
// if unsuccessful, it parses it as a Symbol
func atom(token string) Atom {
	v, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return Symbol(token)
	}
	return Number(v)
}

// tokenize first surrounds all opening and closing parens with
// a space, then returns a slice of strings splitting the input on
// space
func tokenize(s string) []string {
	r := strings.NewReplacer("(", " ( ", ")", " ) ")
	return strings.Fields(r.Replace(s))
}

func main() {
	p := "(begin (define r 10) (* pi (* r r)))"
	fmt.Printf("%#v", tokenize(p))
	k, err := parse(p)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v", k)
}
