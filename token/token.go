package token

const (
	// ops
	plus = iota
	minus
	star
	slash
	greater
	less
	greaterEq
	lessEq
	eq
	eqEq
	notEq
	bitOr
	bitAnd
	bitNot
	logNot
	logAnd
	logOr

	// nonop chars
	arrow
	lparen
	rparen
	lbrace
	rbrace

	// keywords
	let
	mut
	class
	ifK
	elseK
	whileK
	forK
	typeK

	//literals
	intL
	ident
	stringL
	char
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
