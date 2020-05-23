package desugarer

import (
	. "github.com/knarka/yabl/expr"
)

func Desugar(surfaceExpr SurfaceExpr) CoreExpr {
	switch surfaceExpr.Kind {
	case ExprAdd:
		return CAdd(Desugar(surfaceExpr.First()), Desugar(surfaceExpr.Second()))
	case ExprSub:
		return CAdd(Desugar(surfaceExpr.First()), CMult(CNum(-1), Desugar(surfaceExpr.Second())))
	case ExprMult:
		return CMult(Desugar(surfaceExpr.First()), Desugar(surfaceExpr.Second()))
	case ExprCons:
		return CCons(Desugar(surfaceExpr.First()), Desugar(surfaceExpr.Second()))
	case ExprList:
		if len(surfaceExpr.List()) == 0 {
			return CNil()
		} else {
			return CCons(Desugar(surfaceExpr.First()), Desugar(SList(surfaceExpr.List()[1:])))
		}
	case ExprNum:
		return CNum(surfaceExpr.Num())
	case ExprName:
		return CName(surfaceExpr.Name())
	case ExprNil:
		return CNil()
	case ExprTrue:
		return CBool(true)
	case ExprFalse:
		return CBool(false)
	}

	panic("desugar: no such type")
}
