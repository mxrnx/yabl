package parser

import (
	. "github.com/knarka/yabl/expr"
	. "github.com/knarka/yabl/tokenizer"
	"strconv"
)

func parseException(v string) SurfaceExpr {
	panic("parseException: " + v)
}

func expandList(vs []Token) []SurfaceExpr {
	vsm := make([]SurfaceExpr, len(vs))
	for i, v := range vs {
		vsm[i] = Parse(v)
	}
	return vsm
}

func assertLen(list []Token, length int) {
	if len(list) != length {
		panic("incorrect number of arguments to " + list[0].Content)
	}
}

func Parse(t Token) SurfaceExpr {
	switch t.Kind {
	case TokenList:
		if len(t.Children) == 0 {
			return SNil()
		}
		switch t.Children[0].Content {
		case "lst":
			return SList(expandList(t.Children[1:]))
		case "fst":
			assertLen(t.Children, 2)
			return SFst(Parse(t.Children[1]))
		case "snd":
			assertLen(t.Children, 2)
			return SSnd(Parse(t.Children[1]))
		case "+":
			assertLen(t.Children, 3)
			return SAdd(Parse(t.Children[1]), Parse(t.Children[2]))
		case "-":
			assertLen(t.Children, 3)
			return SSub(Parse(t.Children[1]), Parse(t.Children[2]))
		case "*":
			assertLen(t.Children, 3)
			return SMult(Parse(t.Children[1]), Parse(t.Children[2]))
		case "cons":
			assertLen(t.Children, 3)
			return SCons(Parse(t.Children[1]), Parse(t.Children[2]))
		case "fn":
			assertLen(t.Children, 3)
			if t.Children[1].Kind != TokenList {
				panic("second arg to 'fn' should be a list of params")
			}
			params := make([]string, 0)
			for _, el := range t.Children[1].Children {
				p := Parse(el)
				if p.Kind != ExprName {
					panic("function params should be names")
				}
				params = append(params, p.Name())
			}
			for i, el := range params {
				for j := i + 1; j < len(params); j++ {
					if params[j] == el {
						panic("params not unique")
					}
				}
			}
			return SFn(params, Parse(t.Children[2]))
		default:
			args := make([]SurfaceExpr, 0)
			if len(t.Children) > 1 {
				for _, el := range t.Children[1:] {
					p := Parse(el)
					args = append(args, p)
				}
			}
			return SApp(Parse(t.Children[0]), args)
		}

	case TokenSym:
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
	case TokenNum:
		n, err := strconv.Atoi(t.Content)
		if err != nil {
			panic(err)
		}
		return SNum(n)
	}

	return parseException("unparsable token: " + t.Content)
}
