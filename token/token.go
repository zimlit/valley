package token

const (
	// ops
	Plus = iota
	Minus
	Star
	Slash
	Greater
	Less
	GreaterEq
	LessEq
	Eq
	EqEq
	NotEq
	BitOr
	BitAnd
	BitNot
	LogNot
	LogAnd
	LogOr

	// nonop chars
	Arrow
	Lparen
	Rparen
	Lbrace
	Rbrace

	// keywords
	Let
	Mut
	Class
	IfK
	ElseK
	WhileK
	OrK
	TypeK

	//literals
	IntL
	Ident
	StringL
	Char
)

// Token represents a token
type Token struct {
	kind    int
	col     int
	line    int
	literal string
}

// NewToken constructs a token
func NewToken(kind int, col int, line int, literal string) Token {
	return Token{
		kind,
		col,
		line,
		literal,
	}
}

// Line an alias for a slice of tokens
type Line []Token
