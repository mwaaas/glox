package glox

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"glox/token"
	"testing"
)

func TestLexing(t *testing.T) {

	tests := []struct {
		text                 string
		expectedToken        []token.Token
		expectedInvalidToken []token.Token
	}{
		// testing single characters
		{
			`(){},.-;/*+`,
			[]token.Token{
				{
					token.LPAREN,
					"(",
					1,
					2,
				},
				{
					token.RPAREN,
					")",
					1,
					3,
				},
				{
					token.LBRACE,
					"{",
					1,
					4,
				},
				{
					token.RBRACE,
					"}",
					1,
					5,
				},
				{
					token.COMMA,
					",",
					1,
					6,
				},
				{
					token.DOT,
					".",
					1,
					7,
				},
				{
					token.MINUS,
					"-",
					1,
					8,
				},
				{
					token.SEMICOLON,
					";",
					1,
					9,
				},
				{
					token.SLASH,
					`/`,
					1,
					10,
				},
				{
					token.STAR,
					"*",
					1,
					11,
				},
				{
					token.PLUS,
					"+",
					1,
					12,
				},
			},
			[]token.Token{},
		},

		// testing single characters with white spaces
		{
			`(    * + - `,
			[]token.Token{
				{
					token.LPAREN,
					"(",
					1,
					2,
				},
				{
					token.STAR,
					"*",
					1,
					7,
				},
				{
					token.PLUS,
					"+",
					1,
					9,
				},
				{
					token.MINUS,
					"-",
					1,
					11,
				},
			},
			[]token.Token{},
		},

		// testing single characters in multiple line
		{
			`(;
- }`,
			[]token.Token{
				{
					token.LPAREN,
					"(",
					1,
					2,
				},
				{
					token.SEMICOLON,
					";",
					1,
					3,
				},
				{
					token.MINUS,
					"-",
					2,
					2,
				},
				{
					token.RBRACE,
					"}",
					2,
					4,
				},
			},
			[]token.Token{},
		},

		// testing single character with unwanted character
		{
			`( ? )
{?}`,
			[]token.Token{
				{
					token.LPAREN,
					"(",
					1,
					2,
				},
				{
					token.RPAREN,
					")",
					1,
					6,
				},
				{
					token.LBRACE,
					"{",
					2,
					2,
				},
				{
					token.RBRACE,
					"}",
					2,
					4,
				},
			},
			[]token.Token{
				{
					token.ILLEGAL,
					"?",
					1,
					4,
				},
				{
					token.ILLEGAL,
					"?",
					2,
					3,
				},
			},
		},

		// One or two character tokens.
		{
			`( >= {`,
			[]token.Token{
				{
					token.LPAREN,
					"(",
					1,
					2,
				},
				{
					token.GreaterEqual,
					">=",
					1,
					5,
				},
				{
					token.LBRACE,
					"{",
					1,
					7,
				},
			},
			[]token.Token{},
		},

		{
			`> ;`,
			[]token.Token{
				{
					token.GREATER,
					">",
					1,
					2,
				},
				{
					token.SEMICOLON,
					";",
					1,
					4,
				},
			},
			[]token.Token{},
		},

		{
			`+=}`,
			[]token.Token{
				{
					token.PLUS,
					"+",
					1,
					2,
				},
				{
					token.EQUAL,
					"=",
					1,
					3,
				},
				{
					token.RBRACE,
					"}",
					1,
					4,
				},
			},
			[]token.Token{},
		},
		{
			`!+!=}`,
			[]token.Token{
				{
					token.NOT,
					"!",
					1,
					2,
				},
				{
					token.PLUS,
					"+",
					1,
					3,
				},
				{
					token.NOTEQUAL,
					"!=",
					1,
					5,
				},
				{
					token.RBRACE,
					"}",
					1,
					6,
				},
			},
			[]token.Token{},
		},
		{
			`<+<=}`,
			[]token.Token{
				{
					token.LESS,
					"<",
					1,
					2,
				},
				{
					token.PLUS,
					"+",
					1,
					3,
				},
				{
					token.LESSEQUAL,
					"<=",
					1,
					5,
				},
				{
					token.RBRACE,
					"}",
					1,
					6,
				},
			},
			[]token.Token{},
		},

		{
			` // comments are ignored
= // this is ignored too
print "hello world"`,
			[]token.Token{
				{
					token.EQUAL,
					"=",
					2,
					2,
				},
				{
					token.PRINT,
					"print",
					3,
					6,
				},
				{
					token.STRING,
					"\"hello world\"",
					3,
					20,
				},
			},
			[]token.Token{},
		},

		{
			`- text
text2 + "text" 3456"mest !`,
			[]token.Token{
				{
					token.MINUS,
					"-",
					1,
					2,
				},
				{
					token.IDENTIFIER,
					"text",
					1,
					7,
				},
				{
					token.IDENTIFIER,
					"text2",
					2,
					6,
				},
				{
					token.PLUS,
					"+",
					2,
					8,
				},
				{
					token.STRING,
					"\"text\"",
					2,
					15,
				},
				{
					token.NUMBER,
					"3456",
					2,
					20,
				},
			},
			[]token.Token{
				{
					token.ILLEGAL,
					"Unterminated string.",
					2,
					21,
				},
			},
		},
	}

	for index, test := range tests {
		indexText := fmt.Sprintf("index: %d", index)

		acutualToken, invalidToken := lex(test.text)

		assert.Equal(t, test.expectedToken,
			acutualToken, indexText)

		assert.Equal(t, test.expectedInvalidToken,
			invalidToken, indexText)
	}
}
