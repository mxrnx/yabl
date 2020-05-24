package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenizeAddition(t *testing.T) {
	input := "(+ a b)"
	tokens := Tokenize(input)
	if !reflect.DeepEqual(tokens, Token{TokenList, "", []Token{
		Token{TokenSym, "+", nil},
		Token{TokenSym, "a", nil},
		Token{TokenSym, "b", nil},
	}}) {
		t.Errorf("did not tokenize addition properly: got %#v", tokens)
	}
}

func TestTokenizeSingleToken(t *testing.T) {
	input := "chunkybacon   "
	tokens := Tokenize(input)
	if !reflect.DeepEqual(tokens, Token{TokenSym, "chunkybacon", nil}) {
		t.Errorf("did not tokenize single token properly: got %#v", tokens)
	}
}

func TestTokenizeEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on empty input")
		}
	}()

	Tokenize("")
}

func TestTokenizeDifficultCons(t *testing.T) {
	input := "(cons (cons 1 2) (cons 2 3))"
	tokens := Tokenize(input)
	if !reflect.DeepEqual(tokens, Token{TokenList, "", []Token{
		Token{TokenSym, "cons", nil},
		Token{TokenList, "", []Token{
			Token{TokenSym, "cons", nil},
			Token{TokenNum, "1", nil},
			Token{TokenNum, "2", nil},
		}},
		Token{TokenList, "", []Token{
			Token{TokenSym, "cons", nil},
			Token{TokenNum, "2", nil},
			Token{TokenNum, "3", nil},
		}},
	}}) {
		t.Errorf("could not tokenize nested cons: got %#v", tokens)
	}
}

func TestTokenizeList(t *testing.T) {
	tokens := Tokenize("(list 1 2 3)")
	expect := Token{Kind: TokenList, Children: []Token{
		{Kind: TokenSym, Content: "list"},
		{Kind: TokenNum, Content: "1"},
		{Kind: TokenNum, Content: "2"},
		{Kind: TokenNum, Content: "3"},
	}}
	if !reflect.DeepEqual(tokens, expect) {
		t.Errorf("could not tokenize sugared list: expected %#v, but got %#v", expect, tokens)
	}
}

func TestTokenizeListSugar(t *testing.T) {
	tokens := Tokenize("'(1 2 3)")
	expect := Token{Kind: TokenList, Children: []Token{
		{Kind: TokenSym, Content: "lst"},
		{Kind: TokenNum, Content: "1"},
		{Kind: TokenNum, Content: "2"},
		{Kind: TokenNum, Content: "3"},
	}}
	if !reflect.DeepEqual(tokens, expect) {
		t.Errorf("could not tokenize sugared list: expected %#v, but got %#v", expect, tokens)
	}
}

func TestTokenizeMalformedList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on malformed list")
		}
	}()

	Tokenize("())")
}

func TestTokenizeUnmatchedParen(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on malformed list")
		}
	}()

	Tokenize(")")
}
