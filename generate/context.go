package main

import (
	"fmt"
	"io"
	"math/rand"
)

var (
	basicTypes = []string{"-", "int", "uint", "int8", "uint8", "int64", "uint64"}
	idSeq      = make(map[string]int)

	exprDepth    = 2
	numVariables = 5
)

const (
	stmtOas = iota
	stmtReturn
)

type fun struct {
	typ string // returning type
	nam string // name of the function

	lit []lite
	stm []stmt
}

type expr struct {
	typ string // expression type
	str string // string
	atm bool   // if it's atom
}

type lite struct { // a variable
	typ string
	nam string
}

type stmt struct {
	opr int
	lit *lite
	exp *expr
}

func newFun(typ string) *fun {
	if typ == "-" {
		typ = "int"
	}

	f := &fun{
		nam: newId("f"),
		typ: typ,
	}
	f.stm = append(f.stm, stmt{
		opr: stmtReturn,
		exp: f.newExpr(typ, 0),
	})
	return f
}

func (f *fun) dump(w io.Writer) {
	fmt.Fprintf(w, "func %s() %s {\n", f.nam, f.typ)
	fmt.Fprintf(w, "switch {} // prevent inlining\n")
	for i := range f.stm {
		f.stm[i].dump(w)
	}
	fmt.Fprintf(w, "}\n")
}

func (f *fun) newLite(typ string) *lite {

	// placeholder.
	var nl, ns int
	nl, f.lit = len(f.lit), append(f.lit, lite{})
	e := f.newExpr(typ, 0)
	ns, f.stm = len(f.stm), append(f.stm, stmt{})

	f.lit[nl] = lite{
		typ: typ,
		nam: newId("v"),
	}
	f.stm[ns] = stmt{
		opr: stmtOas,
		lit: &f.lit[nl],
		exp: e,
	}
	return &f.lit[nl]
}

func conv(e *expr, typ string) *expr {
	if typ == e.typ {
		return e
	}
	return &expr{
		typ: typ,
		str: fmt.Sprintf("%s(%s)", typ, e.str),
		atm: true,
	}
}

func (f *fun) newExpr(typ string, dep int) *expr {
	if rnd(exprDepth) == 0 || dep > exprDepth*2 {
	loop:
		switch rnd(3) {
		case 0: // constant
			return &expr{
				typ: "-",
				str: fmt.Sprintf("%d", rnd(11)),
				atm: true,
			}

		case 1: // new lit
			if typ == "-" || len(f.lit) >= numVariables {
				break
			}

			l := f.newLite(typ)
			return &expr{
				typ: l.typ,
				str: l.nam,
				atm: true,
			}
		case 2: // reuse lit
			if typ == "-" || len(f.lit) == 0 {
				break
			}

			// Try to pick a variable that has the same type.
			var l *lite
			for i := 0; i < 3 && (l == nil || l.nam == "" || l.typ != typ); i++ {
				l = &f.lit[rnd(len(f.lit))]
			}
			if l.nam == "" {
				// Picked a placeholter.
				break
			}

			return conv(&expr{
				typ: l.typ,
				str: l.nam,
				atm: true,
			}, typ)
		}
		goto loop
	}

retry:
	switch op := choice("+", "-", "*", "()", "<<", ">>", "conv"); op {
	case "+", "-", "*":
		l := f.newExpr(typ, dep+1)
		r := f.newExpr(typ, dep+1)
		if l.typ == "-" && r.typ == "-" {
			typ = "-"
		}

		return &expr{
			typ: typ,
			str: fmt.Sprintf("%s %s %s", l.str, op, r.str),
		}
	case ">>", "<<":
		l := conv(f.newExpr(typ, dep+1), typ)
		r := f.newExpr(choice("-", "uint", "uint8", "uint16", "uint32", "uint64"), dep+1)

		if !r.atm {
			r.str = "(" + r.str + ")"
			r.atm = true
		}

		return &expr{
			typ: typ,
			str: fmt.Sprintf("%s %s %s", l.str, op, r.str),
		}

	case "()":
		e := f.newExpr(typ, dep+1)
		if e.atm {
			return e
		}
		return &expr{
			typ: typ,
			str: fmt.Sprintf("(%s)", e.str),
			atm: true,
		}
	case "conv":
		if typ == "-" || rnd(5) != 0 {
			// Cannot convert to untyped.
			break
		}
		return conv(f.newExpr(choice(basicTypes...), dep+1), typ)
	}
	goto retry
}

func (s *stmt) dump(w io.Writer) {
	switch s.opr {
	case stmtReturn:
		fmt.Fprintf(w, "return %s\n", s.exp.str)
	case stmtOas:
		if s.exp.typ == "-" {
			fmt.Fprintf(w, "%s := %s(%s) // %s\n", s.lit.nam, s.lit.typ, s.exp.str, s.lit.typ)
		} else {
			fmt.Fprintf(w, "%s := %s // %s\n", s.lit.nam, s.exp.str, s.lit.typ)
		}
	}
}

// rnd returns a random integer between 0 and n-1.
func rnd(n int) int {
	return rand.Intn(n)
}

func rndBool() bool {
	return rnd(2) == 0
}

func choice(ch ...string) string {
	return ch[rnd(len(ch))]
}

func newId(prefix string) string {
	idSeq[prefix]++
	return fmt.Sprintf("%v%v", prefix, idSeq[prefix])
}
