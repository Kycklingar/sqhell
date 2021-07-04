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

//func (c *Group) Eval(param *int) string {
//	var ret string
//	for i := 0; i < len(c.conds); i++ {
//		ret += fmt.Sprintf("%s %s\n", c.ops[i], c.conds[i].Eval(param))
//	}
//	return ret
//}

func (c *Group) Eval(param *int) string {
	if len(c.conds) <= 0 {
		return ""
	}

	var ret = c.conds[0].Eval(param)

	for i := 1; i < len(c.conds); i++ {
		ret += fmt.Sprintf("%s %s", c.ops[i], c.conds[i].Eval(param))
	}

	return ret
}
