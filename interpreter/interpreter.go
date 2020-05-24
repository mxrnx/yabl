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

func interpret(coreExpr CoreExpr, env Environment) Value {
	switch coreExpr.Kind {
	case ExprFst:
		return typeCheck(interpret(coreExpr.Expr(), env), ExprCons).Fst()
	case ExprSnd:
		return typeCheck(interpret(coreExpr.Expr(), env), ExprCons).Snd()
	case ExprAdd:
		a := typeCheck(interpret(coreExpr.Head(), env), ExprNum)
		b := typeCheck(interpret(coreExpr.Tail(), env), ExprNum)
		return VNum(a.Num() + b.Num())
	case ExprMult:
		a := typeCheck(interpret(coreExpr.Head(), env), ExprNum)
		b := typeCheck(interpret(coreExpr.Tail(), env), ExprNum)
		return VNum(a.Num() * b.Num())
	case ExprCons:
		return VCons(interpret(coreExpr.Head(), env), interpret(coreExpr.Tail(), env))
	case ExprNil:
		return VNil()
	case ExprNum:
		return VNum(coreExpr.Num())
	case ExprBool:
		return VBool(coreExpr.Bool())
	case ExprName:
		return env.lookup(coreExpr.Name())
	case ExprFn:
		f := coreExpr.Fn()
		return VFn(f.Params, f.Body)
	case ExprApp:
		a := coreExpr.App()
		f := typeCheck(interpret(a.Function, env), ExprFn).Func()
		envNew := Environment{}
		for i, arg := range a.Args {
			envNew = append(envNew, Bind{f.Params[i], interpret(arg, env)})
		}
		return interpret(f.Body, append(envNew, env...))
	}

	return interpException("invalid coreExpr")
}

func Interpret(input string) Value {
	return interpret(desugarer.Desugar(parser.Parse(tokenizer.Tokenize(input))), nil)
}
