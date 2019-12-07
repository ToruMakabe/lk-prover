// %はgoyaccの定義部
%{
package pfparser

import (
	"errors"
	"io"
	"strings"
	"text/scanner"
)

type Expression interface{}

type Token struct {
	Token   int
	Literal string
}

type Literal struct {
	Literal string
}

type NotOpExpr struct {
	Operator rune
	Right    Expression
}

type BinOpExpr struct {
	Left     Expression
	Operator rune
	Right    Expression
}

%}

%union{
  token Token
  expr Expression
}

%type<expr> formula
%type<expr> expr and_expr or_expr not_expr imply_expr paren_expr
%token<token> LITERAL

%left '&' '|' '>'
%right '~'

%%
// 以降はgoyaccの規則部.

formula
  : expr
  {
    $$ = $1
    yylex.(*Lexer).result = $$
  }

expr
	: LITERAL
	{
		$$ = Literal{Literal: $1.Literal}
	}
	| and_expr
	| or_expr
	| imply_expr
	| not_expr
	| paren_expr

and_expr
	: expr '&' expr
	{
		$$ = BinOpExpr{Left: $1, Operator: '&', Right: $3}
	}

or_expr
	: expr '|' expr
	{
		$$ = BinOpExpr{Left: $1, Operator: '|', Right: $3}
	}

imply_expr
	: expr '>' expr
	{
		$$ = BinOpExpr{Left: $1, Operator: '>', Right: $3}
	}

not_expr
	: '~' expr
	{
		$$ = NotOpExpr{Operator: '~', Right: $2}
	}

paren_expr
	: '(' expr ')'
	{
		$$ = $2
	}

%%
// 以降はgoyaccのユーザー定義部. Goで記述する.

// 字句解析器(Lexer)とyaccを用いた構文解析関数(ここから)
type Lexer struct {
	scanner.Scanner
	result Expression
	err	error
}

func (l *Lexer) Lex(lval /* lexer value */ *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Ident {
		token = LITERAL
	}
	lval.token = Token{Token: token, Literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e /* error */ string) {
	l.err = errors.New(e)
}

func Parse(r /* reader */ io.Reader) (Expression, error) {
	l := new(Lexer)
	l.Init(r)
	yyParse(l)
	if l.err != nil {
		return nil, l.err
	}
	return l.result, nil
}
// 字句解析器(Lexer)とyaccを用いた構文解析関数(ここまで)

// PfParseは命題論理式を構文解析し、根に論理結合子があれば [(否定)v1] [論理結合子] [v2]の形式で返す. 論理結合子がなければ [(否定)a]で返す.
func PfParse(pf /* propositional formula */ string) ([]string, error) {
	r := strings.NewReader(pf)
	// yaccで構文木を作成する.
	p, err := Parse(r)
	if err != nil {
		return nil, err
	}

	switch p.(type){
	case BinOpExpr:
		v1 := Eval(p.(BinOpExpr).Left)
		lc := string(rune(p.(BinOpExpr).Operator))
		v2 := Eval(p.(BinOpExpr).Right)
		return []string{v1,lc,v2}, nil
	case NotOpExpr:
		v1 := string(rune(p.(NotOpExpr).Operator)) + Eval(p.(NotOpExpr).Right)
		return []string{v1,"",""}, nil
	case Literal:
		v1 := p.(Literal).Literal
		return []string{v1,"",""}, nil
	}
	return nil, nil
}

// Evalはyaccで作成した構文木を文字列に変換する.
func Eval(e /* expression */ Expression) string {
	switch e.(type) {
	case BinOpExpr:
		left := Eval(e.(BinOpExpr).Left)
		right := Eval(e.(BinOpExpr).Right)
		return "(" + left + string(rune(e.(BinOpExpr).Operator)) + right + ")"
	case NotOpExpr:
		return "(" + string(rune(e.(NotOpExpr).Operator)) + Eval(e.(NotOpExpr).Right)  + ")"
	case Literal:
		return e.(Literal).Literal
	default:
		return ""
	}
}
