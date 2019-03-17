package token

type TokenType string

type Token struct {
	Type         TokenType
	Literal      string
	LineNumber   int
	ColumnNumber int
}

const (

	// single characters
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	SEMICOLON = ";"
	MINUS     = "-"
	STAR      = "*"
	SLASH     = `/`
	DOT       = "."
	PLUS      = "+"

	// one ro tow character token
	GREATER      = ">"
	GreaterEqual = ">="
	EQUAL        = "="
	NOT          = "!"
	NOTEQUAL     = "!="
	LESS         = "<"
	LESSEQUAL    = "<="

	// Literals.
	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// Keywords
	AND    = "and"
	CLASS  = "class"
	ELSE   = "else"
	FALSE  = "false"
	FUN    = "fun"
	FOR    = "for"
	IF     = "if"
	NIL    = "nil"
	OR     = "or"
	PRINT  = "print"
	RETURN = "return"
	SUPER  = "super"
	THIS   = "this"
	TRUE   = "true"
	VAR    = "var"
	WHILE  = "while"

	ILLEGAL = "ILLEGAL"
)

var KEYWORDS = []string{
	AND, CLASS, ELSE, FALSE, FUN, FOR, IF, NIL, OR,
	PRINT, RETURN, SUPER, THIS, TRUE, VAR, WHILE,
}
