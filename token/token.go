package token

const (
	// ops
	Plus = iota
	Minus
	Star
	Slash
	PlusEq
	MinusEq
	StarEq
	SlashEq
	PlusPlus
	MinusMinus
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
	Rshift
	Lshift
	LogNot
	LogAnd
	LogOr
	Dot
	At

	// nonop chars
	Arrow
	FatArrow
	Lparen
	Rparen
	Lbrace
	Rbrace
	Comma
	ColonColon

	// keywords
	Let
	Mut
	Class
	IfK
	ElseK
	WhileK
	ForK
	TypeK
	False
	True

	//literals
	IntL
	FloatL
	Ident
	StringL
	Char
)

// Token represents a token
type Token struct {
	kind    int
	col     int
	line    int
	Literal string
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
