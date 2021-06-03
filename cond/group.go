package cond

import "fmt"

func NewGroup() *Group {
	return new(Group)
}

type Group struct {
	conds []Evaluator
	ops   []string
}

func (c *Group) Add(op string, cond Evaluator) *Group {
	c.conds = append(c.conds, cond)
	c.ops = append(c.ops, op)

	return c
}

func (c *Group) Eval(param *int) string {
	var ret string
	for i := 0; i < len(c.conds); i++ {
		ret += fmt.Sprintf("%s %s\n", c.ops[i], c.conds[i].Eval(param))
	}
	return ret
}
