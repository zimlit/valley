package lexer

import "valley/token"

// Lexer represents the lexer
type Lexer struct {
	currentLine string
	line        int
	pos         int
	source      string
}

// NewLexer constructs a lexer
func NewLexer(source string) Lexer {
	return Lexer{
		"",
		1,
		0,
		source,
	}
}

func (lexer *Lexer) advance() rune {
	lexer.pos++
	if lexer.pos <= len(lexer.source) {
		return runeAt(lexer.source, lexer.pos-1)
	}
	return rune('\000')
}

func (lexer *Lexer) peek() rune {
	if lexer.pos < len(lexer.source) {
		return runeAt(lexer.source, lexer.pos)
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
			lexer.line++
		case rune('+'):
			line = append(line, token.NewToken(token.Plus, lexer.pos, lexer.line, "+"))
		case rune('*'):
			line = append(line, token.NewToken(token.Star, lexer.pos, lexer.line, "*"))
		case rune('/'):
			line = append(line, token.NewToken(token.Slash, lexer.pos, lexer.line, "/"))
		case rune(','):
			line = append(line, token.NewToken(token.Comma, lexer.pos, lexer.line, ","))
		case rune('.'):
			line = append(line, token.NewToken(token.Dot, lexer.pos, lexer.line, "."))
		case rune(';'):
			line = append(line, token.NewToken(token.Semicolon, lexer.pos, lexer.line, ";"))
		case rune('('):
			line = append(line, token.NewToken(token.Lparen, lexer.pos, lexer.line, "("))
		case rune(')'):
			line = append(line, token.NewToken(token.Rparen, lexer.pos, lexer.line, ")"))
		case rune('{'):
			line = append(line, token.NewToken(token.Lbrace, lexer.pos, lexer.line, "{"))
		case rune('}'):
			line = append(line, token.NewToken(token.Rbrace, lexer.pos, lexer.line, "}"))
		case rune('>'):
			if lexer.peek() == '=' {
				line = append(line, token.NewToken(token.GreaterEq, lexer.pos, lexer.line, ">="))
				lexer.pos++
			} else if lexer.peek() == '>' {
				line = append(line, token.NewToken(token.GreaterEq, lexer.pos, lexer.line, ">>"))
				lexer.pos++
			} else {
				line = append(line, token.NewToken(token.Greater, lexer.pos, lexer.line, ">"))
			}
		case rune('<'):
			if lexer.peek() == '=' {
				line = append(line, token.NewToken(token.LessEq, lexer.pos, lexer.line, "<="))
				lexer.pos++
			} else if lexer.peek() == '<' {
				line = append(line, token.NewToken(token.Lshift, lexer.pos, lexer.line, "<<"))
				lexer.pos++
			} else {
				line = append(line, token.NewToken(token.Less, lexer.pos, lexer.line, "<"))
			}
		case rune('-'):
			if lexer.peek() == '>' {
				line = append(line, token.NewToken(token.Arrow, lexer.pos, lexer.line, "->"))
				lexer.pos++
			} else {
				line = append(line, token.NewToken(token.Minus, lexer.pos, lexer.line, "-"))
			}
		}
	}

	return lines, nil
}

func runeAt(str string, idx int) rune {
	runes := []rune(str)
	return runes[idx]
}
