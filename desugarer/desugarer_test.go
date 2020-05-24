package desugarer

import (
	. "github.com/knarka/yabl/expr"
	"github.com/knarka/yabl/parser"
	"github.com/knarka/yabl/tokenizer"
	"reflect"
	"testing"
)

func TestDesugarMult(t *testing.T) {
	expr := Desugar(parser.Parse(tokenizer.Tokenize("(* 2 3)")))
	expect := CMult(CNum(2), CNum(3))
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not desugar multiplication properly: expected %#v, but got %#v", expect, expr)
	}
}

func TestDesugarList(t *testing.T) {
	expr := Desugar(parser.Parse(tokenizer.Tokenize("(lst 1 2 3)")))
	expect := CCons(CNum(1), CCons(CNum(2), CCons(CNum(3), CNil())))
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not desugar list properly: expected %#v, but got %#v", expect, expr)
	}
}

func TestDesugarBool(t *testing.T) {
	expr := Desugar(parser.Parse(tokenizer.Tokenize("(cons true false)")))
	expect := CCons(CBool(true), CBool(false))
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not desugar bools properly: expected %#v, but got %#v", expect, expr)
	}
}
