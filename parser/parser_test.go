package parser

import (
	. "github.com/knarka/yabl/expr"
	. "github.com/knarka/yabl/tokenizer"
	"reflect"
	"testing"
)

func TestParsePlus(t *testing.T) {
	expr := Parse(Tokenize("(+ 3 a)"))
	expect := SAdd(SNum(3), SName("a"))
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse addition properly: expected %#v, but got %#v", expect, expr)
	}
}

func TestParseSub(t *testing.T) {
	expr := Parse(Tokenize("(- 3 a)"))
	expect := SSub(SNum(3), SName("a"))
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse subtraction properly: expected %#v, but got %#v", expect, expr)
	}
}

func TestParseMult(t *testing.T) {
	expr := Parse(Tokenize("(* 3 a)"))
	expect := SMult(SNum(3), SName("a"))
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse multiplication properly: expected %#v, but got %#v", expect, expr)
	}
}

func TestParseNestedCons(t *testing.T) {
	expr := Parse(Tokenize("(cons a (cons 1 2))"))
	expect := SCons(SName("a"), SCons(SNum(1), SNum(2)))
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse nested cons properly: got %#v", expr)
	}
}

func TestParseSingleName(t *testing.T) {
	expr := Parse(Tokenize("x"))
	expect := SName("x")
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse single name right: got %#v", expr)
	}
}

func TestParseUndefinedApplication(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on application to undefined function")
		}
	}()

	Parse(Tokenize("foo bar"))
}

func TestParseList(t *testing.T) {
	expr := Parse(Tokenize("(lst 1 2 3)"))
	expect := SList([]SurfaceExpr{SNum(1), SNum(2), SNum(3)})
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse list right: expected %#v, but got %#v", expect, expr)
	}
}

func TestParseSimpleFn(t *testing.T) {
	expr := Parse(Tokenize("(fn (a) (+ a 3))"))
	expect := SFn([]string{"a"}, SAdd(SName("a"), SNum(3)))
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse simple fn right: expected %#v, but got %#v", expect, expr)
	}
}

func TestParseDoubleParams(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on application to undefined function")
		}
	}()

	Parse(Tokenize("(fn (a a) 7)"))
}
