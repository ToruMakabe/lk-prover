package formula

import (
	"strings"
)

// Eval は命題論理式を評価し、根に論理結合子があれば [(否定)v1] [論理結合子] [v2]の形式で返す. 論理結合子がなければ [(否定)v1]で返す.
func Eval(f /* formula */ string) ([]string, error) {
	r := strings.NewReader(f)
	// goyaccで構文木を作成する.
	p, err := Parse(r)
	if err != nil {
		return nil, err
	}

	switch p.(type) {
	case BinOpExpr:
		v1 := flatten(p.(BinOpExpr).Left)
		lc := string(rune(p.(BinOpExpr).Operator))
		v2 := flatten(p.(BinOpExpr).Right)
		return []string{v1, lc, v2}, nil
	case NotOpExpr:
		v1 := string(rune(p.(NotOpExpr).Operator)) + flatten(p.(NotOpExpr).Right)
		return []string{v1, "", ""}, nil
	case Literal:
		v1 := p.(Literal).Literal
		return []string{v1, "", ""}, nil
	}
	return nil, nil
}

// flattenはgoyaccで作成した構文木を文字列に変換する.
func flatten(e /* expression */ Expression) string {
	switch e.(type) {
	case BinOpExpr:
		left := flatten(e.(BinOpExpr).Left)
		right := flatten(e.(BinOpExpr).Right)
		return "(" + left + string(rune(e.(BinOpExpr).Operator)) + right + ")"
	case NotOpExpr:
		return "(" + string(rune(e.(NotOpExpr).Operator)) + flatten(e.(NotOpExpr).Right) + ")"
	case Literal:
		return e.(Literal).Literal
	default:
		return ""
	}
}
