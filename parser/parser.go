package parser

import (
	. "github.com/knarka/yabl/expr"
	"github.com/knarka/yabl/tokenizer"
	"strconv"
)

func parseException(v string) SurfaceExpr {
	panic("parseException: " + v)
}

func expandList(vs []tokenizer.Token) []SurfaceExpr {
	vsm := make([]SurfaceExpr, len(vs))
	for i, v := range vs {
		vsm[i] = Parse(v)
	}
	return vsm
}

func Parse(t tokenizer.Token) SurfaceExpr {
	switch t.Kind {
	case tokenizer.TokenList:
		if len(t.Children) == 0 {
			return SNil()
		}
		if t.Children[0].Content == "list" {
			return SList(expandList(t.Children[1:]))
		}
		if len(t.Children) == 3 {
			switch t.Children[0].Content {
			case "+":
				return SAdd(Parse(t.Children[1]), Parse(t.Children[2]))
			case "-":
				return SSub(Parse(t.Children[1]), Parse(t.Children[2]))
			case "*":
				return SMult(Parse(t.Children[1]), Parse(t.Children[2]))
			case "cons":
				return SCons(Parse(t.Children[1]), Parse(t.Children[2]))
			default:
				return parseException("illegal first token in list expression: " + t.Children[0].Content)
			}
		}

	case tokenizer.TokenSym:
		switch t.Content {
		case "nil":
			return SNil()
		case "true":
			return STrue()
		case "false":
			return SFalse()
		default:
			return SName(t.Content)
		}
	case tokenizer.TokenNum:
		n, err := strconv.Atoi(t.Content)
		if err != nil {
			panic(err)
		}
		return SNum(n)
	}

	return parseException("unparsable token")
}
