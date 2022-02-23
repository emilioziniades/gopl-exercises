package eval

import (
	"fmt"
)

func (v Var) String() string { return string(v) }

func (l literal) String() string { return fmt.Sprintf("%.6g", float64(l)) }

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	s := c.fn + "("
	switch len(c.args) {
	case 1:
		s += c.args[0].String()
	case 2:
		s += c.args[0].String() + ", " + c.args[1].String()
	}
	s += ")"
	return s
}
