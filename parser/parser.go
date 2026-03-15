package parser

import (
	"fmt"

	"brainrot-lang/lexer"
)

type Parser struct {
	tokens []lexer.Token // all tokens from the lexer
	pos    int           // index of the token we are currently looking at
	errors []string      // collect errors instead of crashing immediately
}

func New(tokens []lexer.Token) *Parser {
	p := &Parser{
		tokens: tokens,
		pos:    0,
		errors: []string{},
	}
	return p
}

// explanation

func (p *Parser) current() lexer.Token { //returns token at current pos
	if p.pos >= len(p.tokens) {
		return lexer.Token{Type: lexer.EOF}
	}
	return p.tokens[p.pos]
}

func (p *Parser) peek() lexer.Token { // looks 1 ahead without moving
	if p.pos+1 >= len(p.tokens) {
		return lexer.Token{Type: lexer.EOF}
	}
	return p.tokens[p.pos+1]
}

func (p *Parser) advance() lexer.Token { // returns current, moves pos forward
	token := p.current()
	if p.pos < len(p.tokens) {
		p.pos++
	}
	return token
}

func (p *Parser) expect(t lexer.TokenType) lexer.Token { // advance IF current matches t, else add error
	if p.current().Type == t {
		return p.advance()
	}
	p.addError(fmt.Sprintf("expected '%s' but got '%s'", t, p.current().Type), p.current().Line)
	return p.current()
}

func (p *Parser) skipNewlines() { // skip NEWLINE tokens between statements
	for p.current().Type == lexer.NEWLINE {
		p.advance()
	}
}

// Errors returns all parser errors collected during Parse()
func (p *Parser) Errors() []string {
	return p.errors
}

// addError records a structured [Skill Issue] error for the parser
func (p *Parser) addError(msg string, line int) {
	p.errors = append(p.errors, fmt.Sprintf(
		"\n[Skill Issue]\nMeh kya ladle meow ghop ghop ghop \n [Parser Error]\n   Line %d, Col 0 → %s\n",
		line, msg,
	))
}

func (p *Parser) Parse() *Program {
	program := &Program{}

	p.skipNewlines()

	for p.current().Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.skipNewlines()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	p.skipNewlines()

	switch p.current().Type {
	case lexer.VAR:
		return p.parseVarStatement()
	case lexer.PRINT:
		return p.parsePrintStatement()
	case lexer.IF:
		return p.parseIfStatement()
	case lexer.WHILE:
		return p.parseWhileStatement()
	case lexer.FOR:
		return p.parseForStatement()
	case lexer.FUNC:
		return p.parseFuncStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	case lexer.BREAK:
		return p.parseBreakStatement()
	case lexer.CONTINUE:
		return p.parseContinueStatement()
	case lexer.IDENT:
		switch p.peek().Type {
    	case lexer.ASSIGN, lexer.PLUS_ASSIGN, lexer.MINUS_ASSIGN,
         	lexer.ASTERISK_ASSIGN, lexer.SLASH_ASSIGN:
        	return p.parseAssignStatement()
    	default:
        	return p.parseExpressionStatement() // handles i++ via expression
    }
	default:
		return p.parseExpressionStatement()
	}
}

