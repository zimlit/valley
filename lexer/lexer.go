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
	if lexer.pos <= len(lexer.source) {
		return runeAt(lexer.source, lexer.pos-1)
	}
	return rune('\000')
}

// Lex returns an slice of lines
func (lexer *Lexer) Lex() ([]token.Line, error) {
	lines := []token.Line{}
	line := []token.Token{}
	for c := lexer.advance(); c != '\000'; c = lexer.advance() {
		switch c {
		case rune('\n'):
			lines = append(lines, line)
			line = []token.Token{}
		}
	}

	return lines, nil
}

func runeAt(str string, idx int) rune {
	runes := []rune(str)
	return runes[idx]
}
