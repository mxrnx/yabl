package expr

type CoreExpr struct {
	Kind    uint8
	Content interface{}
}

func (c CoreExpr) First() CoreExpr {
	t := c.Content.([]CoreExpr)
	return t[0]
}

func (c CoreExpr) Second() CoreExpr {
	t := c.Content.([]CoreExpr)
	return t[1]
}

func (c CoreExpr) Num() int {
	return c.Content.(int)
}

func (c CoreExpr) Bool() bool {
	return c.Content.(bool)
}

func (c CoreExpr) Name() string {
	return c.Content.(string)
}

func CNil() CoreExpr {
	return CoreExpr{ExprNil, nil}
}

func CNum(n int) CoreExpr {
	return CoreExpr{ExprNum, n}
}

func CName(x string) CoreExpr {
	return CoreExpr{ExprName, x}
}

func CBool(b bool) CoreExpr {
	return CoreExpr{ExprBool, b}
}

func CAdd(a, b CoreExpr) CoreExpr {
	return CoreExpr{ExprAdd, []CoreExpr{a, b}}
}

func CMult(a, b CoreExpr) CoreExpr {
	return CoreExpr{ExprMult, []CoreExpr{a, b}}
}

func CCons(a, b CoreExpr) CoreExpr {
	return CoreExpr{ExprCons, []CoreExpr{a, b}}
}
