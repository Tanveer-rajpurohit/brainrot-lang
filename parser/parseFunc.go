package parser

import (
	"brainrot-lang/lexer"
	"brainrot-lang/utils"
)

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


func (p *Parser) parseElseStatement() *BlockStatement {
	p.expect(lexer.ELSE)
	return p.parseBlockStatement()
}


func (p *Parser) parseWhileStatement() *WhileStatement {

	stmt := &WhileStatement{stmtNode: stmtNode{Line: p.current().Line}}

	p.expect(lexer.WHILE)
	stmt.Condition = p.parseExpression()
	stmt.Body = p.parseBlockStatement()

	return stmt

}


func (p *Parser) parseForStatement() *ForStatement {
	stmt := &ForStatement{stmtNode: stmtNode{Line: p.current().Line}}
	p.expect(lexer.FOR)

	// trust_me_bro i = 0; i < 10; i += 1
	if p.current().Type == lexer.VAR { //trust_me_bro i = 0
		stmt.Init = p.parseVarStatement()
	} else if p.current().Type == lexer.IDENT {
		stmt.Init = p.parseAssignStatement() // i = 0
	} else {
		utils.Fatal(utils.NewError("Parser", p.current().Line, p.current().Column, "expected variable declaration or assignment in for loop init"))
	}

	stmt.Condition = p.parseExpression() // i < 10
	stmt.Post = p.parseAssignStatement() // i += 1
	stmt.Body = p.parseBlockStatement()

	return stmt
}


func (p *Parser) parseFuncStatement() *FuncStatement {
	stmt := &FuncStatement{stmtNode: stmtNode{Line: p.current().Line}}

	p.expect(lexer.FUNC)

	nameToken := p.expect(lexer.IDENT)
	stmt.Name = nameToken.Literal

	p.expect(lexer.LPAREN)
	if p.current().Type != lexer.RPAREN {
		stmt.Params = append(stmt.Params, p.expect(lexer.IDENT).Literal)
		for p.current().Type == lexer.COMMA {
			p.expect(lexer.COMMA)
			stmt.Params = append(stmt.Params, p.expect(lexer.IDENT).Literal)
		}
		p.expect(lexer.RPAREN)
	} else {
		p.expect(lexer.RPAREN)
	}

	return stmt
}


func (p *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{stmtNode: stmtNode{Line: p.current().Line}}
	p.expect(lexer.RETURN)
	stmt.Value = p.parseExpression()

	return stmt
}


func (p *Parser) parseBreakStatement() *BreakStatement {
	stmt := &BreakStatement{stmtNode: stmtNode{Line: p.current().Line}}

	p.expect(lexer.BREAK)

	return stmt
}


func (p *Parser) parseContinueStatement() *ContinueStatement {
	stmt := &ContinueStatement{stmtNode: stmtNode{Line: p.current().Line}}

	p.expect(lexer.CONTINUE)

	return stmt
}


func (p *Parser) parseBlockStatement() *BlockStatement {
	stmt := &BlockStatement{stmtNode: stmtNode{Line: p.current().Line}}

	p.expect(lexer.LBRACE)
	
	for p.current().Type != lexer.RBRACE {
		s := p.parseStatement()
		if s != nil {
			stmt.Statements = append(stmt.Statements, s)
		}
		p.skipNewlines()
	}
	
	p.expect(lexer.RBRACE)
	
	return stmt
}