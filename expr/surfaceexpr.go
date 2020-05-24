package expr

type SurfaceExpr struct {
	Kind    uint8
	content interface{}
}

type SurfaceFunction struct {
	Params []string
	Body   SurfaceExpr
}

type SurfaceApplication struct {
	Function SurfaceExpr
	Args     []SurfaceExpr
}

func (s SurfaceExpr) First() SurfaceExpr {
	t := s.content.([]SurfaceExpr)
	return t[0]
}

func (s SurfaceExpr) Second() SurfaceExpr {
	t := s.content.([]SurfaceExpr)
	return t[1]
}

func (s SurfaceExpr) List() []SurfaceExpr {
	return s.content.([]SurfaceExpr)
}

func (s SurfaceExpr) Expr() SurfaceExpr {
	return s.content.(SurfaceExpr)
}

func (s SurfaceExpr) Num() int {
	return s.content.(int)
}

func (s SurfaceExpr) Name() string {
	return s.content.(string)
}

func (s SurfaceExpr) Fn() SurfaceFunction {
	return s.content.(SurfaceFunction)
}

func (s SurfaceExpr) App() SurfaceApplication {
	return s.content.(SurfaceApplication)
}

func SNil() SurfaceExpr {
	return SurfaceExpr{ExprNil, nil}
}

func SNum(n int) SurfaceExpr {
	return SurfaceExpr{ExprNum, n}
}

func SName(x string) SurfaceExpr {
	return SurfaceExpr{ExprName, x}
}

func STrue() SurfaceExpr {
	return SurfaceExpr{ExprTrue, true}
}

func SFalse() SurfaceExpr {
	return SurfaceExpr{ExprFalse, false}
}

func SAdd(a, b SurfaceExpr) SurfaceExpr {
	return SurfaceExpr{ExprAdd, []SurfaceExpr{a, b}}
}

func SSub(a, b SurfaceExpr) SurfaceExpr {
	return SurfaceExpr{ExprSub, []SurfaceExpr{a, b}}
}

func SMult(a, b SurfaceExpr) SurfaceExpr {
	return SurfaceExpr{ExprMult, []SurfaceExpr{a, b}}
}

func SCons(a, b SurfaceExpr) SurfaceExpr {
	return SurfaceExpr{ExprCons, []SurfaceExpr{a, b}}
}

func SList(x []SurfaceExpr) SurfaceExpr {
	return SurfaceExpr{ExprLst, x}
}

func SFst(c SurfaceExpr) SurfaceExpr {
	return SurfaceExpr{ExprFst, c}
}

func SSnd(c SurfaceExpr) SurfaceExpr {
	return SurfaceExpr{ExprSnd, c}
}

func SFn(params []string, body SurfaceExpr) SurfaceExpr {
	return SurfaceExpr{ExprFn, SurfaceFunction{params, body}}
}

func SApp(function SurfaceExpr, args []SurfaceExpr) SurfaceExpr {
	return SurfaceExpr{ExprApp, SurfaceApplication{function, args}}
}
