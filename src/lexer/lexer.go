package lexer

import "github.com/rasulov-emirlan/jazzlang/src/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	currentChar  byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // state position before starting lexing
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			literal := string(ch) + string(l.currentChar)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.currentChar)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case '(':
		tok = newToken(token.LPAREN, l.currentChar)
	case ')':
		tok = newToken(token.RPAREN, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case '{':
		tok = newToken(token.LBRACE, l.currentChar)
	case '}':
		tok = newToken(token.RBRACE, l.currentChar)
	case '-':
		tok = newToken(token.MINUS, l.currentChar)
	case '!':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			literal := string(ch) + string(l.currentChar)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.currentChar)
		}
	case '/':
		tok = newToken(token.SLASH, l.currentChar)
	case '*':
		tok = newToken(token.ASTERISK, l.currentChar)
	case '<':
		tok = newToken(token.LT, l.currentChar)
	case '>':
		tok = newToken(token.GT, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.currentChar) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.currentChar)
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' ||
		'A' <= character && character <= 'Z' || character == '_'
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

// readIdentifier reads a sequence of characters that are letters
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace advances the position until the next non-whitespace character
func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\n' || l.currentChar == '\t' || l.currentChar == '\r' {
		l.readChar()
	}
}

// peekChar returns the next character in the input without advancing the position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
