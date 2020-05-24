package interpreter

import (
	. "github.com/knarka/yabl/expr"
	"reflect"
	"testing"
)

func TestInterpretFstSnd(t *testing.T) {
	expr := interpret(CFst(CSnd(CCons(CNum(1), CCons(CNum(2), CNum(3))))), nil)
	expect := VNum(2)
	if !reflect.DeepEqual(expr, expect) {
		t.Fail()
	}
}

func TestInterpretEnvironment(t *testing.T) {
	expr := interpret(CAdd(CName("foo"), CName("bar")),
		Environment{{"foo", VNum(3)}, {"bar", VNum(10)}})
	expect := VNum(13)
	if !reflect.DeepEqual(expr, expect) {
		t.Fail()
	}
}
