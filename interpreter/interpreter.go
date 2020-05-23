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

func interpret(coreExpr CoreExpr) Value {
	switch coreExpr.Kind {
	case ExprAdd:
		return VNum(interpret(coreExpr.First()).Num() + interpret(coreExpr.Second()).Num())
	case ExprMult:
		return VNum(interpret(coreExpr.First()).Num() * interpret(coreExpr.Second()).Num())
	case ExprCons:
		return VCons(interpret(coreExpr.First()), interpret(coreExpr.Second()))
	case ExprNil:
		return VNil()
	case ExprNum:
		return VNum(coreExpr.Num())
	case ExprBool:
		return VBool(coreExpr.Bool())
	}

	return interpException("invalid coreExpr")
}

func Interpret(input string) Value {
	return interpret(desugarer.Desugar(parser.Parse(tokenizer.Tokenize(input))))
}
