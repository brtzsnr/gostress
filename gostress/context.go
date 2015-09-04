package main

import (
	"fmt"
	"io"
	"math/rand"
)

const (
	exprDepth     = 3
	numVariables  = 10
	numStatements = 15
	numArgs       = 10
)

var (
	// basic expression types. - is untyped integer.
	basicTypes = []string{"bool", "int", "uint", "int8", "uint8", "int64", "uint64"}
	intTypes   = []string{"int", "uint", "int8", "uint8", "int64", "uint64"}
	// maps prefix to next unused ids
	idSeq = make(map[string]int)
	// operator precendence
	precedence = map[string]int{
		"*":  5,
		"/":  5,
		"%":  5,
		"<<": 5,
		">>": 5,
		"&":  5,
		"&^": 5,
		"+":  4,
		"-":  4,
		"|":  4,
		"^":  4,
		"==": 3,
		"!=": 3,
		"<":  3,
		"<=": 3,
		">":  3,
		">=": 3,
		"&&": 2,
		"||": 1,

		"()":    0,
		"call":  0,
		"conv":  0,
		"token": 0,
	}
)

// fnct represents a function.
type fnct struct {
	typ string // returning type
	nam string // name of the function

	arg []*expr // a list of arguments
	lit []*expr // a list of live literals
	stm []*stmt // a list of statements
}

type expr struct {
	typ      string // expression type
	opr      string // operation, e.g. &&, +
	tok      string // token if it's not an expression
	lft, rgt *expr  // operands. right is nil for unary
}

func (e *expr) String() string {
	switch e.opr {
	case "()":
		return fmt.Sprintf("(%s)", e.lft)
	case "call":
		// HACK because argument list is generated early.
		return e.tok
	case "conv":
		return fmt.Sprintf("%s(%s)", e.typ, e.lft)
	case "token":
		return e.tok
	default:
		return fmt.Sprintf("%s %s %s", e.lft, e.opr, prec(e.rgt, "()"))
	}
}

// prec makes e with precence at least as op.
func prec(e *expr, op string) *expr {
	if precedence[e.opr] <= precedence[op] {
		return e
	}
	return &expr{
		typ: e.typ,
		opr: "()",
		lft: e,
	}
}

// conv returns an expression that converts e to typ.
func conv(e *expr, typ string) *expr {
	if typ == e.typ { // already of correct type
		return prec(e, "conv")
	}

	if typ == "bool" {
		return &expr{
			typ: "bool",
			opr: choice("<", "<=", ">", ">=", "!=", "=="),
			lft: e,
			rgt: constant("-"),
		}
	}

	if e.typ == "bool" {
		return conv(&expr{
			typ: "int",
			opr: "token",
			tok: fmt.Sprintf("b2i[%s]", e),
		}, typ)
	}

	return &expr{
		typ: typ,
		opr: "conv",
		lft: e,
	}
}

// constant returns a constant of type typ.
func constant(typ string) *expr {
	if typ == "bool" {
		return &expr{
			typ: "bool",
			opr: "token",
			tok: fmt.Sprint(1 == rnd(2)),
		}
	}

	return &expr{
		typ: "-",
		opr: "token",
		tok: fmt.Sprint(rnd(4)),
	}
}

type stmt struct {
	opr string // operation, e.g :=, _
	lit *expr  // literal
	exp *expr  // expression
}

func newBlock(f *fnct, depth int) {
	quit := false
	nlit := len(f.lit)

	for i := 0; !quit && i < numStatements; i++ {
		op := choice("if", "return", ":=", "=", "+=", "-=", "*=", "<<=", ">>=", "&=", "|=", "&^=", "++", "--")

		switch op {
		case "if":
			if rnd(3) != 0 {
				break
			}
			f.stm = append(f.stm, &stmt{
				opr: "if",
				exp: f.newExpr("bool", 0),
			})
			newBlock(f, depth+1)
			f.stm = append(f.stm, &stmt{
				opr: "}",
			})
		case "return":
			if rnd(numStatements) != 0 {
				break
			}
			f.stm = append(f.stm, &stmt{
				opr: "return",
				exp: f.newExpr(f.typ, 0),
			})
			quit = true
		case ":=":
			l := f.newVariable(choice(basicTypes...))
			f.stm = append(f.stm, &stmt{
				opr: "_",
				lit: l,
			})
		case "=":
			if l := f.getVariable(choice(basicTypes...)); l != nil {
				f.stm = append(f.stm, &stmt{
					opr: op,
					lit: l,
					exp: f.newExpr(l.typ, 0),
				})
			}
		case "+=", "-=", "*=", "&=", "|=", "&^=":
			if l := f.getVariable(choice(intTypes...)); l != nil {
				f.stm = append(f.stm, &stmt{
					opr: op,
					lit: l,
					exp: f.newExpr(l.typ, 0),
				})
			}
		case "<<=", ">>=":
			if l := f.getVariable(choice(intTypes...)); l != nil {
				f.stm = append(f.stm, &stmt{
					opr: op,
					lit: l,
					exp: f.newExpr("uint", 0),
				})
			}
		case "++", "--":
			if l := f.getVariable(choice(intTypes...)); l != nil {
				f.stm = append(f.stm, &stmt{
					opr: op,
					lit: l,
				})
			}
		}
	}

	if depth == 0 && !quit {
		f.stm = append(f.stm, &stmt{
			opr: "return",
			exp: f.newExpr(f.typ, 0),
		})
	}
	f.lit = f.lit[:nlit]
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

	argc := rnd(numArgs + 1)
	for i := 0; i < argc; i++ {
		f.arg = append(f.arg, &expr{
			typ: choice(basicTypes...),
			opr: "token",
			tok: newId("a"),
		})
	}

	newBlock(f, 0)
	return f
}

func (f *fnct) dump(w io.Writer) {
	fmt.Fprintf(w, "func %s(", f.nam)
	for _, a := range f.arg {
		fmt.Fprintf(w, "%s %s, ", a.tok, a.typ)
	}
	fmt.Fprintf(w, ") %s {\n", f.typ)
	fmt.Fprintf(w, "switch {} // prevent inlining\n")
	for i := range f.stm {
		f.stm[i].dump(w)
	}
	fmt.Fprintf(w, "}\n")
}

// newVariable creates a new variable.
func (f *fnct) newVariable(typ string) *expr {
	e := f.newExpr(typ, 0)
	if typ == "-" {
		typ = "int"
	}
	lit := &expr{
		typ: typ,
		opr: "token",
		tok: newId("v"),
	}
	stm := &stmt{
		opr: ":=",
		lit: lit,
		exp: e,
	}
	f.lit = append(f.lit, lit)
	f.stm = append(f.stm, stm)
	return lit
}

// getVariable returns a random live variable. Returns nil if none is found.
// typ == "" means any type.
func (f *fnct) getVariable(typ string) *expr {
	if rnd(10) == 0 {
		return nil
	}
	if rnd(2) == 0 && len(f.lit) > 0 {
		// Pick an already defined literal.
		l := f.lit[rnd(len(f.lit))]
		if l.typ == typ {
			return l
		}
	} else if len(f.arg) > 0 {
		// Pick an argument.
		l := f.arg[rnd(len(f.arg))]
		if l.typ == typ {
			return l
		}
	}
	// Retry.
	return f.getVariable(typ)
}

// getLiteral returns a constant or a variable.
func (f *fnct) getLiteral(typ string) *expr {
	if rnd(50) == 0 { // don't generate many constants
		return constant(typ)
	}
	if l := f.getVariable(typ); l != nil {
		return conv(l, typ)
	}
	return f.getLiteral(typ)
}

// call returns an expression to call a function.
func (f *fnct) call(o *fnct) *expr {
	str := o.nam + "("
	for _, a := range o.arg {
		// HACK because we store the whole function call in tok.
		str += fmt.Sprintf("%s, ", constant(a.typ).tok)
	}
	str += ")"

	return &expr{
		typ: o.typ,
		opr: "call",
		tok: str,
	}
}

func (f *fnct) newExpr(typ string, dep int) *expr {
	if rnd(exprDepth) == 0 || dep > exprDepth {
		return f.getLiteral(typ)
	}

	var ops []string
	if typ == "bool" {
		ops = []string{"&&", "||", "!=", "==", "conv"}
	} else {
		ops = []string{"+", "-", "*" /*"/", "%", */, "&", "|", "^", "&^", "<<", ">>", "++", "--", "conv"}
	}

retry:
	switch op := choice(ops...); op {
	case "&&", "||":
		l := f.newExpr(typ, dep+1)
		r := f.newExpr(typ, dep+1)
		l, r = prec(l, op), prec(r, op)
		return &expr{typ: typ, opr: op, lft: l, rgt: r}

	case "!=", "==":
		ntyp := choice(basicTypes...)
		l := f.newExpr(ntyp, dep+1)
		r := f.newExpr(ntyp, dep+1)
		l, r = prec(l, op), prec(r, op)
		return &expr{typ: typ, opr: op, lft: l, rgt: r}

	case "+", "-", "*", "/", "%", "&", "|", "^":
		l := f.newExpr(typ, dep+1)
		r := f.newExpr(typ, dep+1)
		l, r = prec(l, op), prec(r, op)
		if l.typ == "-" && r.typ == "-" {
			typ = "-"
		}
		return &expr{typ: typ, opr: op, lft: l, rgt: r}
	case ">>", "<<":
		l := f.newExpr(typ, dep+1)
		r := f.newExpr(choice("-", "uint", "uint8", "uint16", "uint32", "uint64"), dep+1)
		if l.typ == "-" {
			l = conv(l, typ)
		}
		l, r = prec(l, op), prec(r, op)
		return &expr{typ: typ, opr: op, lft: l, rgt: r}
	case "conv":
		if typ == "-" || rnd(5) != 0 {
			// Cannot convert to untyped.
			// Also, don't convert very often.
			break
		}
		return conv(f.newExpr(choice(basicTypes...), dep+1), typ)
	}
	goto retry
}

func (s *stmt) dump(w io.Writer) {
	switch s.opr {
	case "if":
		fmt.Fprintf(w, "if %s {\n", s.exp)
	case "}":
		fmt.Fprintf(w, "}\n")
	case "return":
		fmt.Fprintf(w, "return %s\n", s.exp)
	case ":=", "=", "+=", "-=", "*=":
		if s.exp.typ == "-" && s.lit.typ != "int" {
			fmt.Fprintf(w, "%s %s %s(%s) // %s\n", s.lit.tok, s.opr, s.lit.typ, s.exp, s.lit.typ)
		} else {
			fmt.Fprintf(w, "%s %s %s // %s\n", s.lit.tok, s.opr, s.exp, s.lit.typ)
		}
	case "_":
		fmt.Fprintf(w, "_ = %s\n", s.lit.tok)
	case "++", "--":
		fmt.Fprintf(w, "%s%s\n", s.lit.tok, s.opr)
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
