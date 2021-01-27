package lexer

import "valley/token"

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

func (lexer *Lexer) advance() rune {
	lexer.pos++
	lexer.col++
	return runeAt(lexer.source, lexer.pos-1)
}

// Lex returns an slice of lines
func (lexer *Lexer) Lex() ([]token.Line, error) {
	return []token.Line{}, nil
}

func runeAt(str string, idx int) rune {
	runes := []rune(str)
	return runes[idx]
}
