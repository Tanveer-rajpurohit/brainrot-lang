package utils

import (
	"brainrot-lang/lexer"
)

// GetTokenCategory returns the category of a token
func GetTokenCategory(tokType lexer.TokenType) string {
	// Keywords
	if tokType == lexer.VAR || tokType == lexer.IF || tokType == lexer.ELSE ||
		tokType == lexer.WHILE || tokType == lexer.FOR || tokType == lexer.FUNC ||
		tokType == lexer.RETURN || tokType == lexer.PRINT || tokType == lexer.TRUE ||
		tokType == lexer.FALSE || tokType == lexer.NIL || tokType == lexer.BREAK ||
		tokType == lexer.CONTINUE || tokType == lexer.ELSE_IF {
		return "KEYWORD"
	}

	// Literals
	if tokType == lexer.IDENT || tokType == lexer.INT || tokType == lexer.FLOAT || tokType == lexer.STRING {
		return "LITERAL"
	}

	// Operators
	if tokType == lexer.PLUS || tokType == lexer.MINUS || tokType == lexer.ASTERISK || tokType == lexer.SLASH || tokType == lexer.PERCENT || tokType == lexer.POWER || tokType == lexer.INCREMENT || tokType == lexer.DECREMENT {          
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
	if tokType == lexer.ASSIGN || tokType == lexer.PLUS_ASSIGN || tokType == lexer.MINUS_ASSIGN || tokType == lexer.ASTERISK_ASSIGN || tokType == lexer.SLASH_ASSIGN {                                       
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

