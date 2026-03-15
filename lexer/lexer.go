package lexer

import (
	"unicode"
)

// Lexer tokenizes source code
type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // next position to read
	ch           byte // current character
	line         int
	column       int
}

// New creates a new lexer
func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

// readChar advances to the next character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

// peekChar looks at next character without advancing
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// skipWhitespace skips spaces, tabs, etc.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

// skipComment skips a single-line comment starting with #
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

// readNumber reads a number (int or float)
func (l *Lexer) readNumber() (string, TokenType) {
	start := l.position
	tokenType := INT

	// Read integer part
	for unicode.IsDigit(rune(l.ch)) {
		l.readChar()
	}
	// Check for float
	if l.ch == '.' && unicode.IsDigit(rune(l.peekChar())) {
		tokenType = FLOAT
		l.readChar() // consume the '.'
		for unicode.IsDigit(rune(l.ch)) {
			l.readChar()
		}
	}

	return l.input[start:l.position], tokenType

}

// readIdentifier reads an identifier or keyword
func (l *Lexer) readIdentifier() string {
	start := l.position
	for unicode.IsLetter(rune(l.ch)) || unicode.IsDigit(rune(l.ch)) || l.ch == '_' {
		l.readChar()
	}
	return l.input[start:l.position]
}

// readString reads a string literal
func (l *Lexer) readString(quote byte) string {
	l.readChar() // skip the opening quote character
	start := l.position

	result := []byte{}
	for l.ch != quote && l.ch != 0 {
		// Handle escape sequences
		if l.ch == '\\' {
			l.readChar() // consume the backslash
			switch l.ch {
			case 'n':
				result = append(result, '\n')
			case 't':
				result = append(result, '\t')
			case '\\':
				result = append(result, '\\')
			case '"':
				result = append(result, '"')
			case '\'':
				result = append(result, '\'')
			default:
				result = append(result, '\\', l.ch)
			}
		} else {
			result = append(result, l.ch)
		}
		l.readChar()
	}
	_ = start
	l.readChar() // skip the closing quote
	return string(result)
}

// nextToken returns the next token
func (l *Lexer) nextToken() Token {
	l.skipWhitespace()

	line := l.line
	col := l.column

	var tok Token

	switch l.ch {

	case '#':
		l.skipComment()
		return l.nextToken() // skip comment, get next token

	// Arithmetic
	case '+':
		if l.peekChar() == '+' {
			l.readChar()
			tok = Token{INCREMENT, "++", line, col}
		} else if l.peekChar() == '=' {
			l.readChar()
			tok = Token{PLUS_ASSIGN, "+=", line, col}
		} else {
			tok = Token{PLUS, "+", line, col}
		}
		l.readChar()

	case '-':
		if l.peekChar() == '-' {
			l.readChar()
			tok = Token{DECREMENT, "--", line, col}
		} else if l.peekChar() == '=' {
			l.readChar()
			tok = Token{MINUS_ASSIGN, "-=", line, col}
		} else {
			tok = Token{MINUS, "-", line, col}
		}
		l.readChar()

	case '*':
		if l.peekChar() == '*' {
			l.readChar()
			tok = Token{POWER, "**", line, col}
		} else if l.peekChar() == '=' {
			l.readChar()
			tok = Token{ASTERISK_ASSIGN, "*=", line, col}
		} else {
			tok = Token{ASTERISK, "*", line, col}
		}
		l.readChar()

	case '/':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{SLASH_ASSIGN, "/=", line, col}
		} else {
			tok = Token{SLASH, "/", line, col}
		}
		l.readChar()

	case '%':
		tok = Token{PERCENT, "%", line, col}
		l.readChar()

	// Assignment & Comparison
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{EQ, "==", line, col}
		} else if l.peekChar() == '>' {
			l.readChar()
			tok = Token{ARROW, "=>", line, col}
		} else {
			tok = Token{ASSIGN, "=", line, col}
		}
		l.readChar()

	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{NOT_EQ, "!=", line, col}
		} else {
			tok = Token{NOT, "!", line, col}
		}
		l.readChar()

	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{LTE, "<=", line, col}
		} else {
			tok = Token{LT, "<", line, col}
		}
		l.readChar()

	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{GTE, ">=", line, col}
		} else {
			tok = Token{GT, ">", line, col}
		}
		l.readChar()

	// Logical Symbols
	case '&':
		if l.peekChar() == '&' {
			l.readChar()
			tok = Token{AND, "&&", line, col}
			l.readChar()
		} else {
			tok = Token{ILLEGAL, "&", line, col}
			l.readChar()
		}

	case '|':
		if l.peekChar() == '|' {
			l.readChar()
			tok = Token{OR, "||", line, col}
			l.readChar()
		} else {
			tok = Token{ILLEGAL, "|", line, col}
			l.readChar()
		}

	// Delimiters
	case '(':
		tok = Token{LPAREN, "(", line, col}
		l.readChar()
	case ')':
		tok = Token{RPAREN, ")", line, col}
		l.readChar()
	case '{':
		tok = Token{LBRACE, "{", line, col}
		l.readChar()
	case '}':
		tok = Token{RBRACE, "}", line, col}
		l.readChar()
	case '[':
		tok = Token{LBRACKET, "[", line, col}
		l.readChar()
	case ']':
		tok = Token{RBRACKET, "]", line, col}
		l.readChar()

	// Punctuation
	case ';':
		tok = Token{SEMICOLON, ";", line, col}
		l.readChar()
	case ',':
		tok = Token{COMMA, ",", line, col}
		l.readChar()
	case '.':
		tok = Token{DOT, ".", line, col}
		l.readChar()
	case ':':
		tok = Token{COLON, ":", line, col}
		l.readChar()

	// Newline
	case '\n':
		tok = Token{NEWLINE, "\\n", line, col}
		l.readChar()

	// String Literals
	case '"':
		value := l.readString('"')
		tok = Token{STRING, value, line, col}
	case '\'':
		value := l.readString('\'')
		tok = Token{STRING, value, line, col}

	// EOF
	case 0:
		tok = Token{EOF, "", line, col}

	// ── Identifiers, Keywords, Numbers ────────
	default:
		if unicode.IsLetter(rune(l.ch)) || l.ch == '_' {

			literal := l.readIdentifier()
			tokType := LookupIdent(literal)
			return Token{tokType, literal, line, col}

		} else if unicode.IsDigit(rune(l.ch)) {

			literal, tokType := l.readNumber()
			return Token{tokType, literal, line, col}

		} else {

			tok = Token{ILLEGAL, string(l.ch), line, col}
			l.readChar()
		}
	}

	return tok
}

// Tokenize returns all tokens from the input
func (l *Lexer) Tokenize() []Token {
	var tokens []Token
	for {
		tok := l.nextToken()
		tokens = append(tokens, tok)
		if tok.Type == EOF {
			break
		}
	}
	return tokens
}
