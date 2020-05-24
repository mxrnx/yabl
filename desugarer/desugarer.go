package desugarer

import (
	. "github.com/knarka/yabl/expr"
)

func Desugar(surfaceExpr SurfaceExpr) CoreExpr {
	switch surfaceExpr.Kind {
	case ExprFst:
		return CFst(Desugar(surfaceExpr.Expr()))
	case ExprSnd:
		return CSnd(Desugar(surfaceExpr.Expr()))
	case ExprAdd:
		return CAdd(Desugar(surfaceExpr.First()), Desugar(surfaceExpr.Second()))
	case ExprSub:
		return CAdd(Desugar(surfaceExpr.First()), CMult(CNum(-1), Desugar(surfaceExpr.Second())))
	case ExprMult:
		return CMult(Desugar(surfaceExpr.First()), Desugar(surfaceExpr.Second()))
	case ExprCons:
		return CCons(Desugar(surfaceExpr.First()), Desugar(surfaceExpr.Second()))
	case ExprFn:
		f := surfaceExpr.Fn()
		return CFn(f.Params, Desugar(f.Body))
	case ExprApp:
		app := surfaceExpr.App()
		coreArgs := make([]CoreExpr, 0)
		for _, arg := range app.Args {
			coreArgs = append(coreArgs, Desugar(arg))
		}
		return CApp(Desugar(app.Function), coreArgs)
	case ExprLst:
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
