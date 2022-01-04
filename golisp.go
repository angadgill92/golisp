package main

import (
	"fmt"
	"strconv"
	"strings"
)

// tokenize first surrounds all opening and closing parens with
// a space, then returns a slice of strings splitting the input on
// space
func tokenize(s string) []string {
	r := strings.NewReplacer("(", " ( ", ")", " ) ")
	return strings.Fields(r.Replace(s))
}

type Atom interface {
	isAtom() bool
}
type Symbol string

func (s Symbol) isAtom() bool {
	return true
}
type Number float64

func (n Number) isAtom() bool {
	return true
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

func main() {
	input := tokenize(`(add 2 3 4)`)
	fmt.Printf("%#v \n", input)
}
