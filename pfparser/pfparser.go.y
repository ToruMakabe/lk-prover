%{
package pfparser

import (
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

%type<expr> program
%type<expr> expr and_expr or_expr not_expr impl_expr paren_expr
%token<token> LITERAL

%left '&' '|' '>'
%right '~'

%%

program
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
	panic(e)
}

func Parse(r io.Reader) Expression {
	l := new(Lexer)
	l.Init(r)
	yyParse(l)
	return l.result
}

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

func PfParse(pf string) []string {
	r := strings.NewReader(pf)
	p := Parse(r)

	switch p.(type){
	case BinOpExpr:
		a := Eval(p.(BinOpExpr).Left)
		l := string(rune(p.(BinOpExpr).Operator))
		c := Eval(p.(BinOpExpr).Right)
		return []string{a,l,c}
	case NotOpExpr:
		a := string(rune(p.(NotOpExpr).Operator)) + Eval(p.(NotOpExpr).Right)
		return []string{a,"",""}
	case Literal:
		a := p.(Literal).Literal
		return []string{a,"",""}
	}
	return nil
}
