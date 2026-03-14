package parser

import "brainrot-lang/lexer"

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

func (p *Parser) Parse() *Program {}
func (p *Parser) current() lexer.Token {} //returns token at current pos
func (p *Parser) peek() lexer.Token {}   // looks 1 ahead without moving
func (p *Parser) advance() lexer.Token {} // returns current, moves pos forward
func (p *Parser) expect(t lexer.TokenType) lexer.Token {} // advance IF current matches t, else add error
func (p *Parser) skipNewlines() {} // skip NEWLINE tokens between statements