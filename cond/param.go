package cond

import "fmt"

type (
	P string
	N string
	O struct {
		I *int
		S string
	}
)

func (p P) Eval(param *int) string {
	*param++
	return fmt.Sprintf(string(p), *param-1)
}

func (n N) Eval(*int) string {
	return string(n)
}

func (o O) Eval(param *int) string {
	if *o.I == 0 {
		*o.I = *param
		*param++
	}

	return fmt.Sprintf(o.S, *o.I)
}
