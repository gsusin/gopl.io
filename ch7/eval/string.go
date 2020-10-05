package eval

import (
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return b.x.String() + " " + string(b.op) + " " + b.y.String()
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("pow(%s, %s)", c.args[0].String(), c.args[1].String())
	case "sin":
		return fmt.Sprintf("sin(%s)", c.args[0].String())
	case "sqrt":
		return fmt.Sprintf("sqrt(%s)", c.args[0].String())
	default:
		return ""
	}
}

func (e ext) String() string {
	switch e.fn {
	case "max":
		ns := "max("
		for i, n := range e.args {
			if i > 0 {
				ns += ", "
			}
			ns += n.String()
		}
		ns += ")"
		return ns
	default:
		return ""
	}
}
