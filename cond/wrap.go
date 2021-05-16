package cond

type Wrap struct {
	Cond Evaluator
}

func (w Wrap) Eval(param *int) string {
	return "(\n" + w.Cond.Eval(param) + ")\n"
}
