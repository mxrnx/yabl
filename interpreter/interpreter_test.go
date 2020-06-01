package interpreter

import (
	. "github.com/knarka/yabl/expr"
	"reflect"
	"testing"
)

func TestInterpretFstSnd(t *testing.T) {
	expr, _ := interpret(CFst(CSnd(CCons(CNum(1), CCons(CNum(2), CNum(3))))), nil, nil)
	expect := VNum(2)
	if !reflect.DeepEqual(expr, expect) {
		t.Fail()
	}
}

func TestInterpretEnvironment(t *testing.T) {
	expr, _ := interpret(CAdd(CName("foo"), CName("bar")),
		Environment{{"foo", 0}, {"bar", 1}},
		Store{{0, VNum(3)}, {1, VNum(10)}})
	expect := VNum(13)
	if !reflect.DeepEqual(expr, expect) {
		t.Fail()
	}
}
