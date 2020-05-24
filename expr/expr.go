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
	ExprFst
	ExprSnd
	ExprFn
	ExprApp

	// sugar
	ExprSub
	ExprLst
	ExprTrue
	ExprFalse
)
