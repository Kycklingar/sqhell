package cond

import "fmt"

func Group() *group {
	return new(group)
}

type group struct {
	conds []Evaluator
	ops   []string
}

func (c *group) Add(op string, cond Evaluator) *group {
	c.conds = append(c.conds, cond)
	c.ops = append(c.ops, op)

	return c
}

func (c *group) Eval(param *int) string {
	var ret string
	for i := 0; i < len(c.conds); i++ {
		ret += fmt.Sprintf("%s %s\n", c.ops[i], c.conds[i].Eval(param))
	}
	return ret
}
