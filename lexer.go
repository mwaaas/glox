package glox

import (
	"glox/token"
	"strconv"
	"strings"
	"unicode"
)

type lexer struct {
	currentIndex  int
	text          string
	validTokens   []token.Token
	inValidTokens []token.Token
	lineNumber    int
	columnNumber  int
	start         int
}

func newLexer(text string) *lexer {

	return &lexer{0, text,
		[]token.Token{},
		[]token.Token{}, 1, 1, 0}
}

func (l *lexer) getCurrentChar() string {
	return string(l.text[l.currentIndex])
}

func (l *lexer) getNextCharacter() string {
	return string(l.text[l.currentIndex+1])
}
func (l *lexer) getColumnNumber() int {
	return l.columnNumber
}

func (l *lexer) getTokens() ([]token.Token, []token.Token) {

	// parse until the end of the text
	for !l.isEnd() {
		l.start = l.currentIndex
		char := l.advance()
		switch char {

		// single character
		case token.LPAREN:
			l.addDefaultToken(token.LPAREN)
		case token.RPAREN:
			l.addDefaultToken(token.RPAREN)
		case token.LBRACE:
			l.addDefaultToken(token.LBRACE)
		case token.RBRACE:
			l.addDefaultToken(token.RBRACE)
		case token.COMMA:
			l.addDefaultToken(token.COMMA)
		case token.SEMICOLON:
			l.addDefaultToken(token.SEMICOLON)
		case token.MINUS:
			l.addDefaultToken(token.MINUS)
		case token.STAR:
			l.addDefaultToken(token.STAR)
		case token.DOT:
			l.addDefaultToken(token.DOT)
		case token.PLUS:
			l.addDefaultToken(token.PLUS)

		case "\n":
			l.lineNumber += 1
			l.columnNumber = 1

		// one or more tokens
		case token.GREATER:
			if l.match(token.EQUAL) {
				l.addDefaultToken(token.GreaterEqual)
			} else {
				l.addDefaultToken(token.GREATER)
			}

		case token.EQUAL:
			l.addDefaultToken(token.EQUAL)

		case token.NOT:
			if l.match(token.EQUAL) {
				l.addDefaultToken(token.NOTEQUAL)
			} else {
				l.addDefaultToken(token.NOT)
			}

		case token.LESS:
			if l.match(token.EQUAL) {
				l.addDefaultToken(token.LESSEQUAL)
			} else {
				l.addDefaultToken(token.LESS)
			}

		case token.SLASH:
			if l.match(token.SLASH) {
				for l.peek() != "\n" {
					l.advance()
				}
			} else {
				l.addDefaultToken(token.SLASH)
			}
		case `"`:
			l.stringIdentifier()
		default:
			if isNumber(char) {
				l.numberIdentifier()
			} else if isAlphaNumeric(char) {
				l.identifier()
			} else {
				if !strings.ContainsAny(char, " ") {
					l.addInvalidToken(l.getToken(token.ILLEGAL, char))
				}
			}
		}
	}
	return l.validTokens, l.inValidTokens
}

func (l *lexer) peek() string {
	if l.isEnd() {
		return "\n"
	}
	return l.getCurrentChar()
}

/**
Match with the next character
*/
func (l *lexer) match(expected string) bool {
	if l.isEnd() {
		return false
	}
	if l.getCurrentChar() == expected {
		l.advance()
		return true
	}
	return false
}

func (l *lexer) identifier() {
	for isAlphaNumeric(l.peek()) {
		l.advance()
	}

	identifier := l.text[l.start:l.currentIndex]

	// check if its a keyword
	for _, keyword := range token.KEYWORDS {
		if identifier == keyword {
			l.addDefaultToken(token.TokenType(keyword))
			return
		}
	}
	l.addValidToken(
		l.getToken(token.IDENTIFIER, identifier))
}

func (l *lexer) numberIdentifier() {
	for isNumber(l.peek()) {
		l.advance()
	}

	// for decimal numbers
	if l.peek() == "." && !l.isEnd() &&
		isNumber(l.getNextCharacter()) {
		// consume the decimal point
		l.advance()

		for isNumber(l.peek()) {
			l.advance()
		}
	}

	number := l.text[l.start:l.currentIndex]

	l.addValidToken(l.getToken(token.NUMBER, number))
}
func (l *lexer) stringIdentifier() {
	for l.peek() != `"` && !l.isEnd() {
		l.advance()
	}

	if l.isEnd() {
		l.addInvalidToken(token.Token{
			Type:         token.ILLEGAL,
			Literal:      "Unterminated string.",
			LineNumber:   l.lineNumber,
			ColumnNumber: l.columnNumber - (l.currentIndex - l.start - 1)})
	} else {
		l.advance()
		l.addValidToken(
			l.getToken(token.STRING, l.text[l.start:l.currentIndex]))
	}
}

func (l *lexer) getToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{
		Type:         tokenType,
		Literal:      literal,
		LineNumber:   l.lineNumber,
		ColumnNumber: l.getColumnNumber()}
}

func (l *lexer) addDefaultToken(tokenType token.TokenType) {
	l.addToken(token.Token{
		Type:         tokenType,
		Literal:      string(tokenType),
		LineNumber:   l.lineNumber,
		ColumnNumber: l.columnNumber},
		true,
	)
}

/**
 * This adds token and advances to the next character
 */
func (l *lexer) addToken(t token.Token, valid bool) {
	if valid {
		l.validTokens = append(l.validTokens, t)
	} else {
		l.inValidTokens = append(l.inValidTokens, t)
	}
}

/**
 * This adds  invalid token and advances to the next character
 */
func (l *lexer) addInvalidToken(t token.Token) {
	l.addToken(t, false)
}

func (l *lexer) addValidToken(t token.Token) {
	l.addToken(t, true)
}

/**
* This moves to the next character to match
 */
func (l *lexer) advance() string {
	l.currentIndex += 1
	l.columnNumber += 1
	return string(l.text[l.currentIndex-1])
}

/**
Determines when we reach the end of the characters
*/
func (l *lexer) isEnd() bool {
	return len(l.text) == l.currentIndex
}

func isAlphaNumeric(text string) bool {
	char := []rune(text)[0]
	_, err := strconv.Atoi(text)
	if unicode.IsLetter(char) || err == nil {
		return true
	}
	return false
}

func isNumber(text string) bool {
	char := []rune(text)[0]

	if unicode.IsDigit(char) {
		return true
	}
	return false
}

func lex(text string) (validToken []token.Token, invalidToken []token.Token) {
	validToken, invalidToken = newLexer(text).getTokens()
	return
}
