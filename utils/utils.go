package utils

import (
	"fmt"

	"brainrot-lang/lexer"
)

// GetTokenCategory returns the category of a token
func GetTokenCategory(tokType lexer.TokenType) string {
	// Keywords
	if tokType == lexer.VAR || tokType == lexer.IF || tokType == lexer.ELSE ||
		tokType == lexer.WHILE || tokType == lexer.FOR || tokType == lexer.FUNC ||
		tokType == lexer.RETURN || tokType == lexer.PRINT || tokType == lexer.TRUE ||
		tokType == lexer.FALSE || tokType == lexer.NIL || tokType == lexer.BREAK ||
		tokType == lexer.CONTINUE {
		return "KEYWORD"
	}

	// Literals
	if tokType == lexer.IDENT || tokType == lexer.INT || tokType == lexer.FLOAT || tokType == lexer.STRING {
		return "LITERAL"
	}

	// Operators
	if tokType == lexer.PLUS || tokType == lexer.MINUS || tokType == lexer.ASTERISK ||
		tokType == lexer.SLASH || tokType == lexer.PERCENT || tokType == lexer.POWER {
		return "OPERATOR"
	}

	// Comparison
	if tokType == lexer.EQ || tokType == lexer.NOT_EQ || tokType == lexer.LT ||
		tokType == lexer.GT || tokType == lexer.LTE || tokType == lexer.GTE {
		return "COMPARISON"
	}

	// Logical
	if tokType == lexer.AND || tokType == lexer.OR || tokType == lexer.NOT {
		return "LOGICAL"
	}

	// Assignment
	if tokType == lexer.ASSIGN || tokType == lexer.PLUS_ASSIGN || tokType == lexer.MINUS_ASSIGN {
		return "ASSIGNMENT"
	}

	// Delimiters
	if tokType == lexer.LPAREN || tokType == lexer.RPAREN || tokType == lexer.LBRACE ||
		tokType == lexer.RBRACE || tokType == lexer.LBRACKET || tokType == lexer.RBRACKET {
		return "DELIMITER"
	}

	// Punctuation
	if tokType == lexer.SEMICOLON || tokType == lexer.COMMA || tokType == lexer.DOT ||
		tokType == lexer.COLON || tokType == lexer.ARROW {
		return "PUNCTUATION"
	}

	// Special
	if tokType == lexer.EOF {
		return "SPECIAL"
	}

	if tokType == lexer.ILLEGAL {
		return "ILLEGAL"
	}

	return "UNKNOWN"
}

// PrintLexicalTable prints a formatted lexical token table
func PrintLexicalTable(tokens []lexer.Token) {
	// Print header
	fmt.Printf("\n%s┌─────┬──────────────┬──────────────┬─────────────────┬────────┐%s\n", ColorBlue, ColorReset)
	fmt.Printf("%s│ IDX │ TOKEN TYPE   │ LITERAL      │ CATEGORY        │ POS    │%s\n", ColorBlue, ColorReset)
	fmt.Printf("%s├─────┼──────────────┼──────────────┼─────────────────┼────────┤%s\n", ColorBlue, ColorReset)

	// Print each token row
	for i, tok := range tokens {
		if tok.Type == lexer.NEWLINE {
			continue // skip newlines
		}

		// Get token category
		category := GetTokenCategory(tok.Type)

		// Format index
		idx := fmt.Sprintf("[%d]", i)

		// Format token type
		tokType := fmt.Sprintf("%s", tok.Type)

		// Format literal (truncate if too long)
		literal := tok.Literal
		if len(literal) > 12 {
			literal = literal[:9] + "..."
		}
		if tok.Type == lexer.EOF {
			literal = "EOF"
		}

		// Format position
		pos := fmt.Sprintf("L%d:C%d", tok.Line, tok.Column)

		// Print row with colors
		fmt.Printf("%s│ %-3s │ %s%-12s%s │ %s%-12s%s │ %s%-15s%s │ %-6s │%s\n",
			ColorBlue,
			idx,
			ColorCyan, tokType, ColorReset,
			ColorGreen, literal, ColorReset,
			ColorYellow, category, ColorReset,
			pos,
			ColorBlue)
	}

	// Print footer
	fmt.Printf("%s└─────┴──────────────┴──────────────┴─────────────────┴────────┘%s\n", ColorBlue, ColorReset)
	fmt.Printf("\n%sTotal Tokens: %d%s\n\n", ColorBold, len(tokens), ColorReset)
}
