package interpreter

import (
	"github.com/knarka/yabl/desugarer"
	. "github.com/knarka/yabl/expr"
	"github.com/knarka/yabl/parser"
	"github.com/knarka/yabl/tokenizer"
)

func interpException(v string) Value {
	panic("interpException: " + v)
}

func typeCheck(value Value, kind uint8) Value {
	if value.Kind != kind {
		panic("typeCheck failed!")
	} else {
		return value
	}
}

func interpret(coreExpr CoreExpr, env Environment, sto Store) Value {
	switch coreExpr.Kind {
	case ExprFst:
		return typeCheck(interpret(coreExpr.Expr(), env, sto), ExprCons).Fst()
	case ExprSnd:
		return typeCheck(interpret(coreExpr.Expr(), env, sto), ExprCons).Snd()
	case ExprAdd:
		a := typeCheck(interpret(coreExpr.Head(), env, sto), ExprNum)
		b := typeCheck(interpret(coreExpr.Tail(), env, sto), ExprNum)
		return VNum(a.Num() + b.Num())
	case ExprMult:
		a := typeCheck(interpret(coreExpr.Head(), env, sto), ExprNum)
		b := typeCheck(interpret(coreExpr.Tail(), env, sto), ExprNum)
		return VNum(a.Num() * b.Num())
	case ExprCons:
		return VCons(interpret(coreExpr.Head(), env, sto), interpret(coreExpr.Tail(), env, sto))
	case ExprNil:
		return VNil()
	case ExprNum:
		return VNum(coreExpr.Num())
	case ExprBool:
		return VBool(coreExpr.Bool())
	case ExprName:
		return sto.fetch(env.lookup(coreExpr.Name()))
	case ExprFn:
		f := coreExpr.Fn()
		return VFn(f.Params, f.Body)
	case ExprApp:
		a := coreExpr.App()
		f := typeCheck(interpret(a.Function, env, sto), ExprFn).Func()
		envNew := Environment{}
		stoNew := Store{}
		for i, arg := range a.Args {
			loc := newAddr()
			envNew = append(envNew, Pointer{f.Params[i], loc})
			stoNew = append(stoNew, Bind{loc, interpret(arg, env, sto)})

		}
		return interpret(f.Body, append(envNew, env...), append(stoNew, sto...))
	}

	return interpException("invalid coreExpr")
}

func Interpret(input string) Value {
	return interpret(desugarer.Desugar(parser.Parse(tokenizer.Tokenize(input))), nil, nil)
}
