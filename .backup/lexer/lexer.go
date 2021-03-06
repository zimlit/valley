package lexer

import (
	"fmt"
	"unicode"
	"valley/token"
)

// Errors an alias for error handling
type Errors struct {
	Raw []string
}

func (errs Errors) Error() string {
	message := ""
	for _, err := range errs.Raw {
		message += err
	}
	return message
}

type errPos struct {
	col     int
	line    int
	message string
}

func newErrPos(col int, line int, message string) errPos {
	return errPos{
		col,
		line,
		message,
	}
}

// Lexer represents the lexer
type Lexer struct {
	currenttokens string
	line          int
	pos           int
	source        string
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

func (lexer *Lexer) peekNext() rune {
	if !lexer.isAtEnd() {
		return runeAt(lexer.source, lexer.pos+1)
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
		"false": token.False,
		"true":  token.True,
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

func (lexer *Lexer) number() token.Token {
	start := lexer.pos - 1
	t := token.IntL
	for unicode.IsDigit(lexer.peek()) {
		lexer.advance()
		fmt.Println("nd")
	}

	if lexer.peek() == '.' && unicode.IsDigit(lexer.peekNext()) {
		t = token.FloatL
		lexer.advance()
		for unicode.IsDigit(lexer.peek()) {
			lexer.advance()
		}
	}

	return token.NewToken(t, start+1, lexer.line, lexer.source[start:lexer.pos])
}

func (lexer *Lexer) string() (token.Token, errPos) {
	start := lexer.pos
	for lexer.peek() != '"' && !lexer.isAtEnd() {
		if lexer.peek() == '\n' {
			return token.NewToken(token.StringL, start, lexer.line, ""), newErrPos(start, lexer.line, "unterminated string")
		}

		lexer.advance()
	}

	if lexer.isAtEnd() {
		return token.NewToken(token.StringL, start, lexer.line, ""), newErrPos(start, lexer.line, "unterminated string")
	}

	lexer.advance()

	return token.NewToken(token.StringL, start, lexer.line, lexer.source[start:lexer.pos-1]), newErrPos(start, lexer.line, "")
}

// Lex returns an slice of tokenss
func (lexer *Lexer) Lex() ([]token.Token, []string, Errors) {
	lines := []string{}
	tokens := []token.Token{}
	errs := Errors{[]string{}}
	RawErrs := []errPos{}
	start := 0

	for !lexer.isAtEnd() {
		c := lexer.advance()
		switch c {
		// single char tokens
		case rune('\n'):
			lines = append(lines, lexer.source[start:lexer.pos-1])
			for _, err := range RawErrs {
				if err.line == lexer.line {
					message := fmt.Sprint(err.line, " | ", lexer.source[start:lexer.pos])
					for i := range lexer.source {
						if i == err.col {
							message += fmt.Sprint("   ^ ", err.message, "\n\n")
						} else {
							if i < err.col {
								message += " "
							}
						}
					}
					errs.Raw = append(errs.Raw, message)
				}
			}
			start = lexer.pos
			lexer.line++
		case ' ':
		case '\t':
		case '\r':
		case ';':
		case ',':
			tokens = append(tokens, token.NewToken(token.Comma, lexer.pos, lexer.line, ","))
		case '.':
			tokens = append(tokens, token.NewToken(token.Dot, lexer.pos, lexer.line, "."))
		case '(':
			tokens = append(tokens, token.NewToken(token.Lparen, lexer.pos, lexer.line, "("))
		case ')':
			tokens = append(tokens, token.NewToken(token.Rparen, lexer.pos, lexer.line, ")"))
		case '{':
			tokens = append(tokens, token.NewToken(token.Lbrace, lexer.pos, lexer.line, "{"))
		case '}':
			tokens = append(tokens, token.NewToken(token.Rbrace, lexer.pos, lexer.line, "}"))
		case '@':
			tokens = append(tokens, token.NewToken(token.At, lexer.pos, lexer.line, "@"))
		case '~':
			tokens = append(tokens, token.NewToken(token.BitNot, lexer.pos, lexer.line, "~"))
		// two or three char tokens
		case '+':
			if lexer.peek() == '+' {
				tokens = append(tokens, token.NewToken(token.PlusPlus, lexer.pos, lexer.line, "++"))
				lexer.pos++
			} else if lexer.peek() == '=' {
				tokens = append(tokens, token.NewToken(token.PlusEq, lexer.pos, lexer.line, "+="))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.Plus, lexer.pos, lexer.line, "+"))
			}
		case '*':
			if lexer.peek() == '=' {
				tokens = append(tokens, token.NewToken(token.StarEq, lexer.pos, lexer.line, "*="))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.Star, lexer.pos, lexer.line, "*"))
			}
		case '/':
			if lexer.peek() == '=' {
				tokens = append(tokens, token.NewToken(token.SlashEq, lexer.pos, lexer.line, "/="))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.Slash, lexer.pos, lexer.line, "/"))
			}
		case '>':
			if lexer.peek() == '=' {
				tokens = append(tokens, token.NewToken(token.GreaterEq, lexer.pos, lexer.line, ">="))
				lexer.pos++
			} else if lexer.peek() == '>' {
				tokens = append(tokens, token.NewToken(token.GreaterEq, lexer.pos, lexer.line, ">>"))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.Greater, lexer.pos, lexer.line, ">"))
			}
		case '<':
			if lexer.peek() == '=' {
				tokens = append(tokens, token.NewToken(token.LessEq, lexer.pos, lexer.line, "<="))
				lexer.pos++
			} else if lexer.peek() == '<' {
				tokens = append(tokens, token.NewToken(token.Lshift, lexer.pos, lexer.line, "<<"))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.Less, lexer.pos, lexer.line, "<"))
			}
		case '-':
			if lexer.peek() == '>' {
				tokens = append(tokens, token.NewToken(token.Arrow, lexer.pos, lexer.line, "->"))
				lexer.pos++
			} else if lexer.peek() == '-' {
				tokens = append(tokens, token.NewToken(token.MinusMinus, lexer.pos, lexer.line, "--"))
				lexer.pos++
			} else if lexer.peek() == '=' {
				tokens = append(tokens, token.NewToken(token.MinusEq, lexer.pos, lexer.line, "-="))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.Minus, lexer.pos, lexer.line, "-"))
			}
		case '=':
			if lexer.peek() == '=' {
				tokens = append(tokens, token.NewToken(token.EqEq, lexer.pos, lexer.line, "=="))
				lexer.pos++
			} else if lexer.peek() == '>' {
				tokens = append(tokens, token.NewToken(token.FatArrow, lexer.pos, lexer.line, "=>"))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.Eq, lexer.pos, lexer.line, "="))
			}
		case '!':
			if lexer.peek() == '=' {
				tokens = append(tokens, token.NewToken(token.NotEq, lexer.pos, lexer.line, "!="))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.LogNot, lexer.pos, lexer.line, "!"))
			}
		case '|':
			if lexer.peek() == '|' {
				tokens = append(tokens, token.NewToken(token.LogOr, lexer.pos, lexer.line, "||"))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.BitOr, lexer.pos, lexer.line, "|"))
			}
		case '&':
			if lexer.peek() == '&' {
				tokens = append(tokens, token.NewToken(token.LogOr, lexer.pos, lexer.line, "&&"))
				lexer.pos++
			} else {
				tokens = append(tokens, token.NewToken(token.BitAnd, lexer.pos, lexer.line, "&"))
			}
		case ':':
			if lexer.peek() == ':' {
				tokens = append(tokens, token.NewToken(token.ColonColon, lexer.pos, lexer.line, "::"))
				lexer.pos++
			} else {
				RawErrs = append(RawErrs, newErrPos(lexer.pos, lexer.line, "invalid token ':' expected '::'"))
			}
		case '\'':
			var v rune
			start := lexer.pos
			if !lexer.isAtEnd() && lexer.peek() != '\'' && lexer.peek() != '\n' {
				v = lexer.advance()
			}

			if lexer.peek() == '\'' {
				lexer.advance()
				tokens = append(tokens, token.NewToken(token.Char, start, lexer.line, string(v)))
			} else {
				RawErrs = append(RawErrs, newErrPos(start, lexer.line, "unterminated char literal"))
			}
		// other
		case '"':
			tok, err := lexer.string()
			if err.message != "" {
				RawErrs = append(RawErrs, err)
			} else {
				tokens = append(tokens, tok)
			}
		case '#':
			for !lexer.isAtEnd() && lexer.peek() != '\n' {
				lexer.advance()
			}
		default:
			if unicode.IsLetter(c) || c == '_' {
				tokens = append(tokens, lexer.identOrKey())
			} else if unicode.IsDigit(c) {
				tokens = append(tokens, lexer.number())
			} else {
				RawErrs = append(RawErrs, newErrPos(lexer.pos, lexer.line, fmt.Sprint("unexpected char '", string(c), "'")))
			}
		}
	}

	return tokens, lines, errs
}

func runeAt(str string, idx int) rune {
	runes := []rune(str)
	return runes[idx]
}
