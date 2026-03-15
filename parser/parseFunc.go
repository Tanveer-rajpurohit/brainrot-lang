package parser

import "brainrot-lang/lexer"

func (p *Parser) parseVarStatement() *VarStatement {
	stmt := &VarStatement{stmtNode: stmtNode{Line: p.current().Line}}

	p.expect(lexer.VAR)                    // consume "trust_me_bro"
	nameToken := p.expect(lexer.IDENT)     // consume "x", grab the name
	stmt.name = nameToken.Literal

	p.expect(lexer.ASSIGN)                 // consume "="
	stmt.value = p.parseExpression()       // parse whatever is on the right

	return stmt
}


func (p *Parser) parsePrintStatement() *PrintStatement {
	stmt := &PrintStatement{stmtNode: stmtNode{Line : p.current().Line}}

	p.expect(lexer.PRINT)
	stmt.Value = p.parseExpression()

	return stmt
}


func (p *Parser) parseIfStatement() *IfStatement {
	stmt := &IfStatement{stmtNode: stmtNode{Line: p.current().Line}}

	p.expect(lexer.IF)
	stmt.Condition = p.parseExpression()
	stmt.Consequence = p.parseBlockStatement()

	if p.current().Type == lexer.ELSE_IF {
		stmt.ElseIf = p.parseElseIfStatement()
	}

	if p.current().Type == lexer.ELSE {
		stmt.Alternative = p.parseElseStatement()
	}


	return stmt
}

func (p *Parser) parseElseIfStatement() []*ElseIfClause {
	var elseIfs []*ElseIfClause

	for p.current().Type == lexer.ELSE_IF {
		elseIf := &ElseIfClause{stmtNode: stmtNode{Line: p.current().Line}}

		p.expect(lexer.ELSE_IF)

		elseIf.Condition = p.parseExpression()
		elseIf.Body = p.parseBlockStatement()

		elseIfs = append(elseIfs, elseIf)

	}
	return elseIfs
}