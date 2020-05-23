package expr

type SurfaceExpr struct {
	Kind    uint8
	Content interface{}
}

func (s SurfaceExpr) First() SurfaceExpr {
	t := s.Content.([]SurfaceExpr)
	return t[0]
}

func (s SurfaceExpr) Second() SurfaceExpr {
	t := s.Content.([]SurfaceExpr)
	return t[1]
}

func (s SurfaceExpr) List() []SurfaceExpr {
	return s.Content.([]SurfaceExpr)
}

func (s SurfaceExpr) Num() int {
	return s.Content.(int)
}

func (s SurfaceExpr) Name() string {
	return s.Content.(string)
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
	return SurfaceExpr{ExprList, x}
}
