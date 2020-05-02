package desugarer

const (
	PlusExpr      uint8 = 0
	SubExpr       uint8 = 1
	MultExpr      uint8 = 2
	ConsExpr      uint8 = 3
	NameExpr      uint8 = 4
	NumExpr       uint8 = 5
	NilExpr       uint8 = 6
	UndefinedExpr uint8 = 7
)

type CoreExpr struct {
	Kind uint8
	Name string
	Num  int
	Args []CoreExpr
}
