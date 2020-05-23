package expr

const (
	// core
	ExprNil = iota
	ExprNum
	ExprName
	ExprBool
	ExprAdd
	ExprMult
	ExprCons

	// sugar
	ExprSub
	ExprList
	ExprTrue
	ExprFalse
)
