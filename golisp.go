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

func main() {
	input := tokenize(`(add 2 3 4)`)
	fmt.Printf("%#v \n", input)
}
