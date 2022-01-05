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

type List []Atom

func (l List) expression() {}

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
	k := []Atom{
		atom(`2.000`),
		atom("hello"),
		atom("\"hello\""),
	}
	fmt.Printf("%#v type: %T\n", k, k)
}
