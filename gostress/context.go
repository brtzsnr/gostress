package main

import (
	"fmt"
	"io"
	"math/rand"
)

var (
	basicTypes = []string{"int", "uint", "int8", "uint8", "int64", "uint64"}
	idSeq      = make(map[string]int)

	exprDepth     = 5
	numVariables  = 20
	numStatements = 20
	numArgs       = 5
)

// fnct represents a function.
type fnct struct {
	typ string // returning type
	nam string // name of the function

	arg []vrbl
	lit []vrbl // a list of live literals
	stm []stmt // a list of statements
}

type expr struct {
	typ string // expression type
	str string // string
	atm bool   // if it's atom
}

type vrbl struct { // a variable
	typ string // variable's type
	nam string // variable's name.
}

type stmt struct {
	opr string // operation, e.g :=, _
	lit *vrbl  // literal
	exp *expr  // expression
}

// newFnct generates and returns a new function.
func newFnct(typ string) *fnct {
	if typ == "-" {
		typ = "int"
	}
	f := &fnct{
		nam: newId("f") + "_ssa",
		typ: typ,
	}

	argc := rnd(numArgs+1)
	for i := 0; i < argc; i++ {
		f.arg = append(f.arg, vrbl{
			nam: newId("a"),
			typ: choice(basicTypes...),
		})
	}

	for i := 0; i < numStatements; i++ {
		switch o := choice(":=", "=", "++", "--"); o {
		case "return":
			f.stm = append(f.stm, stmt{
				opr: "return",
				exp: f.newExpr(typ, 0),
			})
		case ":=":
			l := f.newVariable(choice(basicTypes...))
			f.stm = append(f.stm, stmt{
				opr: "_",
				lit: l,
			})
		case "=":
			if l := f.getVariable(""); l != nil {
				f.stm = append(f.stm, stmt{
					opr: o,
					lit: l,
					exp: f.newExpr(l.typ, 0),
				})
			}
		case "++", "--":
			if l := f.getVariable(""); l != nil {
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
	fmt.Fprintf(w, "func %s(", f.nam)
	for _, a := range f.arg {
		fmt.Fprintf(w, "%s %s, ", a.nam, a.typ)
	}
	fmt.Fprintf(w, ") %s {\n", f.typ)
	fmt.Fprintf(w, "switch {} // prevent inlining\n")
	for i := range f.stm {
		f.stm[i].dump(w)
	}
	fmt.Fprintf(w, "}\n")
}

func (f *fnct) newVariable(typ string) *vrbl {
	// placeholder.
	var nl, ns int
	nl, f.lit = len(f.lit), append(f.lit, vrbl{})
	e := f.newExpr(typ, 0)
	ns, f.stm = len(f.stm), append(f.stm, stmt{})

	if typ == "-" {
		typ = "int"
	}
	f.lit[nl] = vrbl{
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

// getVariable returns a random live variable. Returns nil if none is found.
// typ == "" means any type.
func (f *fnct) getVariable(typ string) *vrbl {
	if rnd(10) == 0 {
		return nil
	}
	if rnd(2) == 0 && len(f.lit) > 0 {
		l := &f.lit[rnd(len(f.lit))]
		if typ == "" || l.typ == typ {
			return l
		}
	} else if len(f.arg) > 0 {
		l := &f.arg[rnd(len(f.arg))]
		if typ == "" || l.typ == typ {
			return l
		}
	}
	return f.getVariable(typ)
}

// getLiteral returns a constant or a variable.
func (f *fnct) getLiteral(typ string) *expr {
	if rnd(25) == 0 { // don't generate many constants
		return constant()
	}
	if l := f.getVariable(typ); l != nil {
		return conv(&expr{
			typ: l.typ,
			str: l.nam,
			atm: true,
		}, typ)
	}
	return f.getLiteral(typ)
}

func (f *fnct) call(o *fnct) *expr {
	str := o.nam + "("
	for range o.arg {
		str += fmt.Sprintf("%d, ", rnd(10))
	}
	str += ")"

	return &expr{
		typ: o.typ,
		str: str,
	}
}

func constant() *expr {
	return &expr{
		typ: "-",
		str: fmt.Sprintf("%d", rnd(4)),
		atm: true,
	}
}


func (f *fnct) newExpr(typ string, dep int) *expr {
	if rnd(exprDepth) == 0 || dep > exprDepth {
		return f.getLiteral(typ)
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
	case ":=", "=":
		if s.exp.typ == "-" && s.lit.typ != "int" {
			fmt.Fprintf(w, "%s %s %s(%s) // %s\n", s.lit.nam, s.opr, s.lit.typ, s.exp.str, s.lit.typ)
		} else {
			fmt.Fprintf(w, "%s %s %s // %s\n", s.lit.nam, s.opr, s.exp.str, s.lit.typ)
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
