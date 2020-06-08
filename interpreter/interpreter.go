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

func interpret(coreExpr CoreExpr, env Environment, sto Store) (Value, Store) {
	switch coreExpr.Kind {
	case ExprFst:
		valA, stoA := interpret(coreExpr.Expr(), env, sto)
		return valA.Fst(), stoA
	case ExprSnd:
		valA, stoA := interpret(coreExpr.Expr(), env, sto)
		return valA.Snd(), stoA
	case ExprAdd:
		valA, stoA := interpret(coreExpr.Head(), env, sto)
		valB, stoB := interpret(coreExpr.Tail(), env, stoA)
		return VNum(valA.Num() + valB.Num()), stoB
	case ExprMult:
		valA, stoA := interpret(coreExpr.Head(), env, sto)
		valB, stoB := interpret(coreExpr.Tail(), env, stoA)
		return VNum(valA.Num() * valB.Num()), stoB
	case ExprCons:
		valA, stoA := interpret(coreExpr.Head(), env, sto)
		valB, stoB := interpret(coreExpr.Tail(), env, stoA)
		return VCons(valA, valB), stoB
	case ExprNil:
		return VNil(), sto
	case ExprNum:
		return VNum(coreExpr.Num()), sto
	case ExprBool:
		return VBool(coreExpr.Bool()), sto
	case ExprName:
		return sto.fetch(env.lookup(coreExpr.Name())), sto
	case ExprFn:
		f := coreExpr.Fn()
		return VFn(f.Params, f.Body), sto
	case ExprApp:
		a := coreExpr.App()
		valF, _ := interpret(a.Function, env, sto)
		f := valF.Func()
		envNew := Environment{}
		stoNew := Store{}
		for i, arg := range a.Args {
			loc := newAddr()
			envNew = append(envNew, Pointer{f.Params[i], loc})
			valArg, _ := interpret(arg, env, sto)
			stoNew = append(stoNew, Bind{loc, valArg})

		}
		return interpret(f.Body, append(envNew, env...), append(stoNew, sto...))
	}

	return interpException("invalid coreExpr"), nil
}

func Interpret(input string) Value {
	value, _ := interpret(desugarer.Desugar(parser.Parse(tokenizer.Tokenize(input))), nil, nil)
	return value
}
