package parser

import (
	"fmt"

	"brainrot-lang/lexer"
)

func (p *Parser) parseVarStatement() *VarStatement {
	stmt := &VarStatement{stmtNode: stmtNode{Line: p.current().Line}}

	p.expect(lexer.VAR)                    // consume "trust_me_bro"
	nameToken := p.expect(lexer.IDENT)     // consume "x", grab the name
	stmt.Name = nameToken.Literal

	p.expect(lexer.ASSIGN)                 // consume "="
	stmt.Value = p.parseExpression()       // parse whatever is on the right

	return stmt
}


func (p *Parser) parsePrintStatement() *PrintStatement {
	stmt := &PrintStatement{stmtNode: stmtNode{Line : p.current().Line}}

	p.expect(lexer.PRINT)
	p.expect(lexer.LPAREN)
	stmt.Value = p.parseExpression()
	p.expect(lexer.RPAREN)

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
		p.errors = append(p.errors, fmt.Sprintf(
            "[Skill Issue] \nexpected variable declaration or assignment but got '%s' at line %d",
			p.current().Literal, p.current().Line,
		))
	}

	p.expect(lexer.SEMICOLON) 
	stmt.Condition = p.parseExpression() // i < 10
	p.expect(lexer.SEMICOLON)  
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

	stmt.Body = p.parseBlockStatement()

	return stmt
}


func (p *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{stmtNode: stmtNode{Line: p.current().Line}}
	p.expect(lexer.RETURN)
	if p.current().Type == lexer.NEWLINE || p.current().Type == lexer.RBRACE || p.current().Type == lexer.EOF {
    	stmt.Value = nil  // bare return
	} else {
    	stmt.Value = p.parseExpression()
	}

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
	
	for p.current().Type != lexer.RBRACE && p.current().Type != lexer.EOF {
		s := p.parseStatement()
		if s != nil {
			stmt.Statements = append(stmt.Statements, s)
		}
		p.skipNewlines()
	}

	p.expect(lexer.RBRACE)

	return stmt
}


func (p *Parser) parseAssignStatement() *AssignStatement {
	stmt := &AssignStatement{stmtNode: stmtNode{Line: p.current().Line}}

	stmt.Name =  p.expect(lexer.IDENT).Literal

	switch p.current().Type {
    case lexer.ASSIGN, lexer.PLUS_ASSIGN, lexer.MINUS_ASSIGN,
         lexer.ASTERISK_ASSIGN, lexer.SLASH_ASSIGN:
        stmt.Operator = p.current().Literal
        p.advance()
		
    default:
        p.errors = append(p.errors, fmt.Sprintf(
            "[Skill Issue] \nexpected assignment operator but got '%s' at line %d",
            p.current().Literal, p.current().Line,
        ))
    }

	stmt.Value = p.parseExpression()

	return stmt
}


func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{stmtNode: stmtNode{Line: p.current().Line}}

	stmt.Value = p.parseExpression()

	return stmt
}