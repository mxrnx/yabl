package parser

import (
	"github.com/knarka/yabl/tokenizer"
	"reflect"
	"testing"
)

func TestParsePlus(t *testing.T) {
	expr := Parse(tokenizer.Tokenize("(+ 3 a)"))
	expect := SurfaceExpr{Kind: PlusExpr, Args: []SurfaceExpr{{Kind: NumExpr, Num: 3}, {Kind: NameExpr, Name: "a"}}}
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse addition properly: expected %#v, but got %#v", expect, expr)
	}
}

func TestParseSub(t *testing.T) {
	expr := Parse(tokenizer.Tokenize("(- 3 a)"))
	expect := SurfaceExpr{Kind: SubExpr, Args: []SurfaceExpr{{Kind: NumExpr, Num: 3}, {Kind: NameExpr, Name: "a"}}}
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse subtraction properly: expected %#v, but got %#v", expect, expr)
	}
}

func TestParseMult(t *testing.T) {
	expr := Parse(tokenizer.Tokenize("(* 3 a)"))
	expect := SurfaceExpr{Kind: MultExpr, Args: []SurfaceExpr{{Kind: NumExpr, Num: 3}, {Kind: NameExpr, Name: "a"}}}
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse multiplication properly: expected %#v, but got %#v", expect, expr)
	}
}

func TestParseNestedCons(t *testing.T) {
	expr := Parse(tokenizer.Tokenize("(cons a (cons 1 2))"))
	if !reflect.DeepEqual(expr, SurfaceExpr{Kind: ConsExpr, Args: []SurfaceExpr{
		{Kind: NameExpr, Name: "a"},
		{Kind: ConsExpr, Args: []SurfaceExpr{
			{Kind: NumExpr, Num: 1},
			{Kind: NumExpr, Num: 2},
		}}}}) {
		t.Errorf("did not parse nested cons properly: got %#v", expr)
	}
}

func TestParseSingleName(t *testing.T) {
	expr := Parse(tokenizer.Tokenize("x"))
	if !reflect.DeepEqual(expr, SurfaceExpr{Kind: NameExpr, Name: "x"}) {
		t.Errorf("did not parse single name right: got %#v", expr)
	}
}

func TestParseUndefinedApplication(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on application to undefined function")
		}
	}()

	Parse(tokenizer.Tokenize("foo bar"))
}

func TestParseList(t *testing.T) {
	expr := Parse(tokenizer.Tokenize("(list 1 2 3)"))
	expect := SurfaceExpr{Kind: ListExpr, Args: []SurfaceExpr{
		{Kind: NumExpr, Num: 1},
		{Kind: NumExpr, Num: 2},
		{Kind: NumExpr, Num: 3},
	}}
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not parse list right: expected %#v, but got %#v", expect, expr)
	}
}
