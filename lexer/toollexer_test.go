package lexer

import (
	"fmt"
	"testing"

	"toolip-go/token"
)

func TestNextToken(t *testing.T) {
	input := `. .. = * - / // ** % ++ -- > < << >> <> -> ..= aba "caba" 1231321312.12321;`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.DOT, "."},
		{token.CONCAT, ".."},
		{token.ASSIGN, "="},
		{token.MULT, "*"},
		{token.MINUS, "-"},
		{token.DIV, "/"},
		{token.FDIV, "//"},
		{token.EXPO, "**"},
		{token.MOD, "%"},
		{token.INCREMENT, "++"},
		{token.DECREMENT, "--"},
		{token.GT, ">"},
		{token.LT, "<"},
		{token.LSHIFT, "<<"},
		{token.RSHIFT, ">>"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.ARROW, "->"},
		{token.CONCATEQ, "..="},
		{token.IDENTIFIER, "aba"},
		{token.STRVAL, "\"caba\""},
		{token.FLOATVAL, "1231321312.123213"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		i++
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("lexer/toollexer_test.go:%d:%d. expected=(%q: %q), got=(%q: %q).", l.lineNum, l.linePosition, tt.expectedType, tt.expectedValue, tok.Type, tok.Value)
		}
		fmt.Printf("Token(%q, %q)\n", tok.Type, tok.Value)
	}
}
