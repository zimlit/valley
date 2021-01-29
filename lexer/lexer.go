package lexer

import (
	"unicode"
	"valley/token"
)

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
	return runeAt(lexer.source, lexer.pos-1)
}

func (lexer *Lexer) isAtEnd() bool {
	return lexer.pos >= len(lexer.source)
}

func (lexer *Lexer) peek() rune {
	if !lexer.isAtEnd() {
		return runeAt(lexer.source, lexer.pos)
	}
	return rune('\000')
}

func (lexer *Lexer) identOrKey() token.Token {
	keywords := map[string]int{
		"let":   token.Let,
		"mut":   token.Mut,
		"class": token.Class,
		"if":    token.IfK,
		"else":  token.ElseK,
		"while": token.WhileK,
		"for":   token.ForK,
		"type":  token.TypeK,
	}
	start := lexer.pos - 1
	for unicode.IsLetter(lexer.peek()) || unicode.IsDigit(lexer.peek()) || lexer.peek() == '_' {
		lexer.advance()
	}
	text := lexer.source[start:lexer.pos]
	t := keywords[text]
	if t == 0 {
		t = token.Ident
	}
	return token.NewToken(t, start, lexer.line, text)
}

// Lex returns an slice of lines
func (lexer *Lexer) Lex() ([]token.Line, error) {
	lines := []token.Line{}
	line := []token.Token{}
	for !lexer.isAtEnd() {
		c := lexer.advance()
		switch c {
		// single char tokens
		case rune('\n'):
			lines = append(lines, line)
			line = []token.Token{}
			lexer.line++
		case rune(' '):
		case rune('\t'):
		case rune('\r'):
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
		case rune('@'):
			line = append(line, token.NewToken(token.At, lexer.pos, lexer.line, "@"))
		// two char tokens
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
		case rune('='):
			if lexer.peek() == '=' {
				line = append(line, token.NewToken(token.EqEq, lexer.pos, lexer.line, "=="))
				lexer.pos++
			} else if lexer.peek() == '>' {
				line = append(line, token.NewToken(token.FatArrow, lexer.pos, lexer.line, "=>"))
				lexer.pos++
			} else {
				line = append(line, token.NewToken(token.Eq, lexer.pos, lexer.line, "="))
			}
		case rune('!'):
			if lexer.peek() == '=' {
				line = append(line, token.NewToken(token.NotEq, lexer.pos, lexer.line, "!="))
				lexer.pos++
			} else {
				line = append(line, token.NewToken(token.LogNot, lexer.pos, lexer.line, "!"))
			}
		case rune('|'):
			if lexer.peek() == '|' {
				line = append(line, token.NewToken(token.LogOr, lexer.pos, lexer.line, "||"))
				lexer.pos++
			} else {
				line = append(line, token.NewToken(token.BitOr, lexer.pos, lexer.line, "|"))
			}
		case rune('&'):
			if lexer.peek() == '&' {
				line = append(line, token.NewToken(token.LogOr, lexer.pos, lexer.line, "&&"))
				lexer.pos++
			} else {
				line = append(line, token.NewToken(token.BitAnd, lexer.pos, lexer.line, "&"))
			}
		case rune('~'):
			line = append(line, token.NewToken(token.BitNot, lexer.pos, lexer.line, "~"))
			// other
		default:
			if unicode.IsLetter(c) || c == '_' {
				line = append(line, lexer.identOrKey())
			}
		}
	}

	return lines, nil
}

func runeAt(str string, idx int) rune {
	runes := []rune(str)
	return runes[idx]
}
