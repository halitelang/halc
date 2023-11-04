package lexer

import "strings"

type TokenType string

func (l *Lexer) Lex() []Token {
	var tokens []Token
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}
	return tokens
}

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
	MODULO    TokenType = "MODULO"
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
	LBRACKET  TokenType = "LBRACKET"
	RBRACKET  TokenType = "RBRACKET"
	VARIABLE  TokenType = "VARIABLE"
)

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"var":    VARIABLE,
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

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = Token{Type: EQ, Value: literal}
		} else {
			tok = newToken(ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '[':
		tok = newToken(LBRACKET, l.ch)
	case ']':
		tok = newToken(RBRACKET, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: NOT_EQ, Value: string(ch) + string(l.ch)}
		} else {
			tok = newToken(BANG, l.ch)
		}
	case '/':
		if l.peekChar() == '/' {
			// Consume the comment and move to the end of the line
			l.readChar() // Consume the second '/'
			for l.ch != '\n' && l.ch != 0 {
				l.readChar()
			}

			if l.ch == 0 {
				tok.Value = ""
				tok.Type = EOF
			}
		} else {
			tok = newToken(SLASH, l.ch)
		}
	case '*':
		tok = newToken(ASTERISK, l.ch)
	case '<':
		tok = newToken(LT, l.ch)
	case '>':
		tok = newToken(GT, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case '%':
		tok = newToken(MODULO, l.ch)
	case 0:
		tok.Value = ""
		tok.Type = EOF
	case '"':
		tok.Type = STRING
		tok.Value = l.readString()
	default:
		if isLetter(l.ch) {
			tok.Value = l.readIdentifier()
			tok.Type = LookupIdent(tok.Value)
			return tok
		} else if isDigit(l.ch) {
			tok.Value = l.readNumber()
			// Determine if the token is INT or FLOAT
			if strings.Contains(tok.Value, ".") {
				tok.Type = FLOAT
			} else {
				tok.Type = INT
			}
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	hasDecimal := false // New flag to indicate a decimal point

	for isDigit(l.ch) || (l.ch == '.' && isDigit(l.peekChar())) {
		if l.ch == '.' {
			if hasDecimal { // Second decimal point encountered, break out
				break
			}
			hasDecimal = true
		}
		l.readChar()
	}

	// If the number includes a decimal point, it is a FLOAT
	if hasDecimal {
		return l.input[position:l.position]
	}
	// Else, it is an INT
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
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
