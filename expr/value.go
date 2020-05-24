package expr

import (
	"strconv"
)

type Value struct {
	Kind    uint8
	Content interface{}
}

type ValueFunction struct {
	Params []string
	Body   CoreExpr
}

func (v Value) Num() int {
	return v.Content.(int)
}

func (v Value) Fst() Value {
	return v.Content.([]Value)[0]
}

func (v Value) Snd() Value {
	return v.Content.([]Value)[1]
}

func (v Value) Func() ValueFunction {
	return v.Content.(ValueFunction)
}

func (v Value) Pretty() string {
	switch v.Kind {
	case ExprNil:
		return "nil"
	case ExprNum:
		return strconv.Itoa(v.Content.(int))
	case ExprBool:
		return strconv.FormatBool(v.Content.(bool))
	case ExprCons:
		t := v.Content.([]Value)
		return "(" + t[0].Pretty() + ", " + t[1].Pretty() + ")"
	case ExprFn:
		t := v.Content.(ValueFunction)
		return "<fn: " + strconv.Itoa(len(t.Params)) + " Params>"
	default:
		panic("Could not prettify Value of unknown type: " + strconv.Itoa((int(v.Kind))))
	}
}

func VNil() Value {
	return Value{ExprNil, nil}
}

func VNum(n int) Value {
	return Value{ExprNum, n}
}

func VBool(b bool) Value {
	return Value{ExprBool, b}
}

func VCons(a Value, b Value) Value {
	return Value{ExprCons, []Value{a, b}}
}

func VFn(params []string, body CoreExpr) Value {
	return Value{ExprFn, ValueFunction{params, body}}
}
