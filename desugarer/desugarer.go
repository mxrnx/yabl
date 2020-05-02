package desugarer

import (
	"github.com/knarka/yabl/parser"
)

func Desugar(expr parser.SurfaceExpr) CoreExpr {
	switch expr.Kind {
	case parser.PlusExpr:
		return CoreExpr{Kind: PlusExpr, Args: []CoreExpr{Desugar(expr.Args[0]), Desugar(expr.Args[1])}}
	case parser.SubExpr:
		right := Desugar(expr.Args[1])
		right.Num *= -1 // assume type checker guarantees this is a NumExpr
		return CoreExpr{Kind: PlusExpr, Args: []CoreExpr{Desugar(expr.Args[0]), right}}
	case parser.MultExpr:
		return CoreExpr{Kind: MultExpr, Args: []CoreExpr{Desugar(expr.Args[0]), Desugar(expr.Args[1])}}
	case parser.ConsExpr:
		return CoreExpr{Kind: ConsExpr, Args: []CoreExpr{Desugar(expr.Args[0]), Desugar(expr.Args[1])}}
	case parser.ListExpr:
		if len(expr.Args) == 0 {
			return CoreExpr{Kind: NilExpr}
		} else {
			return CoreExpr{Kind: ConsExpr, Args: []CoreExpr{Desugar(expr.Args[0]), Desugar(parser.SurfaceExpr{Kind: parser.ListExpr, Args: expr.Args[1:]})}}
		}
	case parser.NumExpr:
		return CoreExpr{Kind: NumExpr, Num: expr.Num}
	case parser.NameExpr:
		return CoreExpr{Kind: NameExpr, Name: expr.Name}
	}
	return CoreExpr{Kind: UndefinedExpr}
}
