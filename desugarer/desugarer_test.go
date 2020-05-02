package desugarer

import (
	"github.com/knarka/yabl/parser"
	"github.com/knarka/yabl/tokenizer"
	"reflect"
	"testing"
)

func TestDesugarMult(t *testing.T) {
	expr := Desugar(parser.Parse(tokenizer.Tokenize("(* 2 3)")))
	expect := CoreExpr{Kind: MultExpr, Args: []CoreExpr{
		{Kind: NumExpr, Num: 2},
		{Kind: NumExpr, Num: 3},
	}}
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not desugar multiplication properly: expected %#v, but got %#v", expect, expr)
	}
}

func TestDesugarList(t *testing.T) {
	expr := Desugar(parser.Parse(tokenizer.Tokenize("(list 1 2 3)")))
	expect := CoreExpr{Kind: ConsExpr, Args: []CoreExpr{
		{Kind: NumExpr, Num: 1},
		{Kind: ConsExpr, Args: []CoreExpr{
			{Kind: NumExpr, Num: 2},
			{Kind: ConsExpr, Args: []CoreExpr{
				{Kind: NumExpr, Num: 3},
				{Kind: NilExpr},
			}},
		}},
	}}
	if !reflect.DeepEqual(expr, expect) {
		t.Errorf("did not desugar list properly: expected %#v, but got %#v", expect, expr)
	}
}