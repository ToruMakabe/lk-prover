// %はgoyaccのヘッダ定義
%{
package pfparser

import (
	"fmt"
	"io"
	"os"
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
%type<expr> expr and_expr or_expr not_expr impl_expr paren_expr
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
	| impl_expr
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

impl_expr
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

const inputFormatMsg = "Please input LK sequent as (assumptions) |- (conclutions)\nNagation:~, And:&, Or:|, Implication:>\nYou can specify multiple assumtions/conclutions delimitted by comma\nSample: A&B,C |- A,B\n"

// 字句解析器(Lexer)とyaccを用いた構文解析処理(ここから)
type Lexer struct {
	scanner.Scanner
	result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Ident {
		token = LITERAL
	}
	lval.token = Token{Token: token, Literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	fmt.Println()
	fmt.Println("Syntax error!!")
	fmt.Println()
	fmt.Println(inputFormatMsg)
	os.Exit(1)
//	panic(e)
}

func Parse(r io.Reader) Expression {
	l := new(Lexer)
	l.Init(r)
	yyParse(l)
	return l.result
}
// 字句解析器(Lexer)とyaccを用いた構文解析処理(ここまで)

// Evalはyaccで作成した構文木を文字列に変換する.
func Eval(e Expression) string {
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

// PfParseは命題論理式を構文解析し、根に論理結合子があれば [(否定)v1] [論理結合子] [v2]の形式で返す. 論理結合子がなければ [(否定)a]で返す.
func PfParse(pf string) []string {
	r := strings.NewReader(pf)
	// yaccで構文木を作成する.
	p := Parse(r)

	switch p.(type){
	case BinOpExpr:
		v1 := Eval(p.(BinOpExpr).Left)
		lc := string(rune(p.(BinOpExpr).Operator))
		v2 := Eval(p.(BinOpExpr).Right)
		return []string{v1,lc,v2}
	case NotOpExpr:
		v1 := string(rune(p.(NotOpExpr).Operator)) + Eval(p.(NotOpExpr).Right)
		return []string{v1,"",""}
	case Literal:
		v1 := p.(Literal).Literal
		return []string{v1,"",""}
	}
	return nil
}
