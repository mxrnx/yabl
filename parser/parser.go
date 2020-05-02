package parser

import (
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
			return SurfaceExpr{Kind: NilExpr}
		}
		if len(t.Children) == 3 {
			switch t.Children[0].Content {
			case "+":
				return SurfaceExpr{Kind: PlusExpr, Args: []SurfaceExpr{Parse(t.Children[1]), Parse(t.Children[2])}}
			case "-":
				return SurfaceExpr{Kind: SubExpr, Args: []SurfaceExpr{Parse(t.Children[1]), Parse(t.Children[2])}}
			case "*":
				return SurfaceExpr{Kind: MultExpr, Args: []SurfaceExpr{Parse(t.Children[1]), Parse(t.Children[2])}}
			case "list":
				return SurfaceExpr{Kind: ListExpr, Args: expandList(t.Children[1:])}
			case "cons":
				return SurfaceExpr{Kind: ConsExpr, Args: []SurfaceExpr{Parse(t.Children[1]), Parse(t.Children[2])}}
			}
		}

	case tokenizer.TokenSym:
		return SurfaceExpr{Kind: NameExpr, Name: t.Content}
	case tokenizer.TokenNum:
		c, err := strconv.Atoi(t.Content)
		if err != nil {
			panic(err)
		}
		return SurfaceExpr{Kind: NumExpr, Num: c}
	}

	return parseException("unparsable token: " + t.Content)
}
