package main

import (
	"fmt"
	"math"
	"testing"
)

func TestExprStrings(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		//!+Eval
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))

		printed := fmt.Sprint(expr)
		exprPrint, err := Parse(printed)
		if err != nil {
			t.Error(err)
		}
		gotPrint := fmt.Sprintf("%.6g", exprPrint.Eval(test.env))

		fmt.Println(test.expr, "==>", printed)

		if got != gotPrint {
			t.Errorf("\tPre-print gives %s, post-print gives %s", got, gotPrint)
		}
	}
}
