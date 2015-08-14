package main

import (
	"fmt"
	"io"
	"math/rand"
)

var (
	basicTypes = []string{"-", "int", "uint", "int8", "uint8", "int64", "uint64"}
	idSeq      = make(map[string]int)

	exprDepth     = 5
	numVariables  = 10
	numStatements = 10
)

// fnct represents a function.
type fnct struct {
	typ string // returning type
	nam string // name of the function

	lit []lite // a list of live literals
	stm []stmt // a list of statements
}

type expr struct {
	typ string // expression type
	str string // string
	atm bool   // if it's atom
}

type lite struct { // a variable
	typ string // variable's type
	nam string // variable's name.
}

type stmt struct {
	opr string // operation, e.g :=, _
	lit *lite  // literal
	exp *expr  // expression
}

// newFun generates and returns a new function.
func newFun(typ string) *fnct {
	if typ == "-" {
		typ = "int"
	}
	f := &fnct{
		nam: newId("f") + "_ssa",
		typ: typ,
	}

	for i := 0; i < numStatements; i++ {
		switch o := choice(":=", "++", "--"); o {
		case "return":
			f.stm = append(f.stm, stmt{
				opr: "return",
				exp: f.newExpr(typ, 0),
			})
		case ":=":
			l := f.newLite(choice(basicTypes...))
			f.stm = append(f.stm, stmt{
				opr: "_",
				lit: l,
			})
		case "++", "--":
			if l := f.getLit(""); l != nil {
				f.stm = append(f.stm, stmt{
					opr: o,
					lit: l,
				})
			}
		}
	}

	f.stm = append(f.stm, stmt{
		opr: "return",
		exp: f.newExpr(typ, 0),
	})
	return f
}

func (f *fnct) dump(w io.Writer) {
	fmt.Fprintf(w, "func %s() %s {\n", f.nam, f.typ)
	fmt.Fprintf(w, "switch {} // prevent inlining\n")
	for i := range f.stm {
		f.stm[i].dump(w)
	}
	fmt.Fprintf(w, "}\n")
}

func (f *fnct) newLite(typ string) *lite {
	// placeholder.
	var nl, ns int
	nl, f.lit = len(f.lit), append(f.lit, lite{})
	e := f.newExpr(typ, 0)
	ns, f.stm = len(f.stm), append(f.stm, stmt{})

	if typ == "-" {
		typ = "int"
	}
	f.lit[nl] = lite{
		typ: typ,
		nam: newId("v"),
	}
	f.stm[ns] = stmt{
		opr: ":=",
		lit: &f.lit[nl],
		exp: e,
	}
	return &f.lit[nl]
}

// conv returns an expression that converts e to typ.
func conv(e *expr, typ string) *expr {
	if typ == e.typ && e.atm { // already atom and of correct type
		return e
	}
	if typ == e.typ { // correcty type, but not an atom
		return &expr{
			typ: typ,
			str: fmt.Sprintf("(%s)", e.str),
			atm: true,
		}
	}
	return &expr{
		typ: typ,
		str: fmt.Sprintf("%s(%s)", typ, e.str),
		atm: true,
	}
}

// getLit returns a random live literal. Returns nil if none is found.
func (f *fnct) getLit(typ string) *lite {
	if len(f.lit) == 0 {
		return nil
	}
	for i := 0; i < 5; i++ {
		l := &f.lit[rnd(len(f.lit))]
		if (typ == "" || l.typ == typ) && l.nam != "" {
			return l
		}
	}
	return nil
}

func (f *fnct) newExpr(typ string, dep int) *expr {
	if rnd(exprDepth) == 0 || dep > exprDepth {
	loop:
		switch rnd(3) {
		case 0: // constant
			if rnd(25) != 0 { // don't generate many constants
				break
			}
			return &expr{
				typ: "-",
				str: fmt.Sprintf("%d", rnd(4)),
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
			if l := f.getLit(typ); l != nil {
				return conv(&expr{
					typ: l.typ,
					str: l.nam,
					atm: true,
				}, typ)
			}
		}
		goto loop
	}

retry:
	switch op := choice("+", "-", "*", "&", "|", "^", "()", "<<", ">>", "++", "--", "conv"); op {
	case "+", "-", "*", "&", "|", "^":
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
		if e.typ == "-" {
			typ = "-"
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
	case "return":
		fmt.Fprintf(w, "return %s\n", s.exp.str)
	case ":=":
		if s.exp.typ == "-" && s.lit.typ != "int" {
			fmt.Fprintf(w, "%s := %s(%s) // %s\n", s.lit.nam, s.lit.typ, s.exp.str, s.lit.typ)
		} else {
			fmt.Fprintf(w, "%s := %s // %s\n", s.lit.nam, s.exp.str, s.lit.typ)
		}
	case "_":
		fmt.Fprintf(w, "_ = %s\n", s.lit.nam)
	case "++", "--":
		fmt.Fprintf(w, "%s%s\n", s.lit.nam, s.opr)
	}
}

// rnd returns a random integer between 0 and n-1.
func rnd(n int) int {
	return rand.Intn(n)
}

func rndBool() bool {
	return rnd(2) == 0
}

// choice returns a random element from ch.
func choice(ch ...string) string {
	return ch[rnd(len(ch))]
}

func newId(prefix string) string {
	idSeq[prefix]++
	return fmt.Sprintf("%v%v", prefix, idSeq[prefix])
}
