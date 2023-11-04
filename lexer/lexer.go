package lexer

import "strings"

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF       TokenType = "EOF"
	IDENT     TokenType = "IDENT"
	INT       TokenType = "INT"
	STRING    TokenType = "STRING"
	FLOAT     TokenType = "FLOAT"
	ASSIGN    TokenType = "ASSIGN"
	PLUS      TokenType = "PLUS"
	MINUS     TokenType = "MINUS"
	BANG      TokenType = "BANG"
	ASTERISK  TokenType = "ASTERISK"
	SLASH     TokenType = "SLASH"
	LT        TokenType = "LT"
	GT        TokenType = "GT"
	EQ        TokenType = "EQ"
	NOT_EQ    TokenType = "NOT_EQ"
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	LBRACE    TokenType = "LBRACE"
	RBRACE    TokenType = "RBRACE"
	FUNCTION  TokenType = "FUNCTION"
	LET       TokenType = "LET"
	TRUE      TokenType = "TRUE"
	FALSE     TokenType = "FALSE"
	IF        TokenType = "IF"
	ELSE      TokenType = "ELSE"
	RETURN    TokenType = "RETURN"
)

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (t *Lexer) readChar() {
	if t.readPosition >= len(t.input) {
		t.ch = 0
	} else {
		t.ch = t.input[t.readPosition]
	}
	t.position = t.readPosition
	t.readPosition++
}

func (t *Lexer) NextToken() Token {
	var tok Token

	t.skipWhitespace()

	switch t.ch {
	case '=':
		if t.peekChar() == '=' {
			ch := t.ch
			t.readChar()
			literal := string(ch) + string(t.ch)
			tok = Token{Type: EQ, Value: literal}
		} else {
			tok = newToken(ASSIGN, t.ch)
		}
	case ';':
		tok = newToken(SEMICOLON, t.ch)
	case '(':
		tok = newToken(LPAREN, t.ch)
	case ')':
		tok = newToken(RPAREN, t.ch)
	case ',':
		tok = newToken(COMMA, t.ch)
	case '+':
		tok = newToken(PLUS, t.ch)
	case '-':
		tok = newToken(MINUS, t.ch)
	case '!':
		if t.peekChar() == '=' {
			ch := t.ch
			t.readChar()
			tok = Token{Type: NOT_EQ, Value: string(ch) + string(t.ch)}
		} else {
			tok = newToken(BANG, t.ch)
		}
	case '/':
		if t.peekChar() == '/' {
			// Consume the comment and move to the end of the line
			t.readChar() // Consume the second '/'
			for t.ch != '\n' && t.ch != 0 {
				t.readChar()
			}

			if t.ch == 0 {
				tok.Value = ""
				tok.Type = EOF
			}
		} else {
			tok = newToken(SLASH, t.ch)
		}
	case '*':
		tok = newToken(ASTERISK, t.ch)
	case '<':
		tok = newToken(LT, t.ch)
	case '>':
		tok = newToken(GT, t.ch)
	case '{':
		tok = newToken(LBRACE, t.ch)
	case '}':
		tok = newToken(RBRACE, t.ch)
	case 0:
		tok.Value = ""
		tok.Type = EOF
	case '"':
		tok.Type = STRING
		tok.Value = t.readString()
	default:
		if isLetter(t.ch) {
			tok.Value = t.readIdentifier()
			tok.Type = LookupIdent(tok.Value)
			return tok
		} else if isDigit(t.ch) {
			tok.Value = t.readNumber()
			// Determine if the token is INT or FLOAT
			if strings.Contains(tok.Value, ".") {
				tok.Type = FLOAT
			} else {
				tok.Type = INT
			}
			return tok
		} else {
			tok = newToken(ILLEGAL, t.ch)
		}
	}

	t.readChar()
	return tok
}

func (t *Lexer) readIdentifier() string {
	position := t.position
	for isLetter(t.ch) {
		t.readChar()
	}
	return t.input[position:t.position]
}

func (t *Lexer) readString() string {
	position := t.position + 1
	for {
		t.readChar()
		if t.ch == '"' || t.ch == 0 {
			break
		}
	}
	return t.input[position:t.position]
}

func (t *Lexer) readNumber() string {
	position := t.position
	hasDecimal := false // New flag to indicate a decimal point

	for isDigit(t.ch) || (t.ch == '.' && isDigit(t.peekChar())) {
		if t.ch == '.' {
			if hasDecimal { // Second decimal point encountered, break out
				break
			}
			hasDecimal = true
		}
		t.readChar()
	}

	// If the number includes a decimal point, it is a FLOAT
	if hasDecimal {
		return t.input[position:t.position]
	}
	// Else, it is an INT
	return t.input[position:t.position]
}

func (t *Lexer) peekChar() byte {
	if t.readPosition >= len(t.input) {
		return 0
	} else {
		return t.input[t.readPosition]
	}
}

func (t *Lexer) skipWhitespace() {
	for t.ch == ' ' || t.ch == '\t' || t.ch == '\n' || t.ch == '\r' {
		t.readChar()
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Value: string(ch)}
}
