package main

import (
	"fmt"
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
func main() {
	input := tokenize(`(add 2 3 4)`)
	fmt.Printf("%#v \n", input)
}
