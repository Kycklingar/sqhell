package cond

import (
	"fmt"
	"testing"
)

func TestGroup(t *testing.T) {
	var o = new(int)

	g1 := NewGroup().
		Add(",", P("$%d")).
		Add(",", P("$%d")).
		Add(",", O{S: "$%d", I: o}).
		Add(",", O{S: "$%d", I: o})

	fmt.Println(g1.Eval(new(int)))

	g2 := NewGroup().
		Add("\nAND", P("field = $%d")).
		Add("\nAND", P("type = $%d")).
		Add("\nOR", N("tag LIKE 'test'")).
		Add(
			"\nAND",
			Wrap{NewGroup().
				Add("\nAND", P("var = $%d")).
				Add("\nAND", P("id = $%d")),
			},
		)

	fmt.Println(g2.Eval(new(int)))
}
