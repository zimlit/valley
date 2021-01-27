package lexer

// Lexer represents the lexer
type Lexer struct {
	currentLine string
	line        int
	col         int
	pos         int
	source      string
}

// NewLexer constructs a lexer
func NewLexer(source string) Lexer {
	return Lexer{
		"",
		1,
		1,
		0,
		source,
	}
}
