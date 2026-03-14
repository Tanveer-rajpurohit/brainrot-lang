package utils

import "fmt"

// BRLError represents a compiler error with location info
type BRLError struct {
    Line    int
    Column  int
    Phase   string // "Lexer", "Parser", "Interpreter"
    Message string
}

func (e *BRLError) Error() string {
    return fmt.Sprintf(
        "\nMeh kya ladle meow ghop ghop ghop \n [%s Error]\n   Line %d, Col %d → %s\n",
        e.Phase, e.Line, e.Column, e.Message,
    )
}


// NewError creates a new BRLError
func NewError(phase string, line, col int, msg string) *BRLError {
    return &BRLError{
        Phase:   phase,
        Line:    line,
        Column:  col,
        Message: msg,
    }
}