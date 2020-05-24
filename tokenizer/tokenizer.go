package tokenizer

import (
	"regexp"
)

type Token struct {
	Kind     uint8
	Content  string  // number or symbol
	Children []Token // for lists
}

const (
	TokenSym  uint8 = 0 // letter symbols
	TokenNum  uint8 = 1 // numbers
	TokenList uint8 = 2 // all nesting structures
)

func tokenizeException(v string) {
	panic("tokenizeException: " + v)
}

func removeEmptyTokens(elements []string) []string {
	i := 0
	for _, x := range elements {
		if !regexp.MustCompile(`^\s*$`).MatchString(x) {
			elements[i] = x
			i++
		}
	}
	return elements[:i]
}

func standardizeTokens(elements []string) ([]string, Token) {
	var add Token
	var token Token
	var children []Token

	if len(elements) == 0 {
		tokenizeException("unexpected EOF")
	}
	el := elements[0]
	elements = elements[1:]

	switch el {
	case "(":
		for elements[0] != ")" {
			elements, add = standardizeTokens(elements)
			children = append(children, add)

		}
		elements = elements[1:]
		token = Token{TokenList, "", children}
	case ")":
		tokenizeException("unmatched ')'")
	default:
		if regexp.MustCompile(`^\d+$`).MatchString(el) {
			token = Token{TokenNum, el, nil}
		} else {
			token = Token{TokenSym, el, nil}
		}
	}
	return elements, token
}

func Tokenize(input string) Token {
	input = string(regexp.MustCompile(`'\(`).ReplaceAll([]byte(input), []byte("(lst ")))
	input = string(regexp.MustCompile(`\(`).ReplaceAll([]byte(input), []byte("( ")))
	input = string(regexp.MustCompile(`\)`).ReplaceAll([]byte(input), []byte(" )")))
	elements := regexp.MustCompile(`\s`).Split(input, -1)

	e, out := standardizeTokens(removeEmptyTokens(elements))
	if len(e) != 0 {
		tokenizeException("trailing elements: ")
	}
	return out
}
