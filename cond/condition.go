package cond

type Evaluator interface {
	Eval(*int) string
}
