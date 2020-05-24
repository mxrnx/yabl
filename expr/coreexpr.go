package expr

type CoreExpr struct {
	Kind    uint8
	content interface{}
}

type CoreFunction struct {
	Params []string
	Body   CoreExpr
}

type CoreApplication struct {
	Function CoreExpr
	Args     []CoreExpr
}

func (c CoreExpr) Head() CoreExpr {
	t := c.content.([]CoreExpr)
	return t[0]
}

func (c CoreExpr) Tail() CoreExpr {
	t := c.content.([]CoreExpr)
	return t[1]
}

func (c CoreExpr) Expr() CoreExpr {
	return c.content.(CoreExpr)
}

func (c CoreExpr) Num() int {
	return c.content.(int)
}

func (c CoreExpr) Bool() bool {
	return c.content.(bool)
}

func (c CoreExpr) Name() string {
	return c.content.(string)
}

func (c CoreExpr) Fn() CoreFunction {
	return c.content.(CoreFunction)
}

func (c CoreExpr) App() CoreApplication {
	return c.content.(CoreApplication)
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

func CFst(c CoreExpr) CoreExpr {
	return CoreExpr{ExprFst, c}
}

func CSnd(c CoreExpr) CoreExpr {
	return CoreExpr{ExprSnd, c}
}

func CFn(params []string, body CoreExpr) CoreExpr {
	return CoreExpr{ExprFn, CoreFunction{params, body}}
}

func CApp(function CoreExpr, args []CoreExpr) CoreExpr {
	return CoreExpr{ExprApp, CoreApplication{function, args}}
}
