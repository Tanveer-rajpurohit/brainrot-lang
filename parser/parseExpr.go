package parser

import (
	"fmt"
	"strconv"

	"brainrot-lang/lexer"
)

// EXPRESSION CHAIN — bottom to top by precedence
//
// Each function calls the one BELOW it.
// Lower in the chain = higher precedence = evaluates first.
//
// parseExpression          lowest  (entry point)
//   └── parseOr            ||
//         └── parseAnd     &&
//               └── parseEquality    == !=
//                     └── parseComparison  < > <= >=
//                           └── parseTerm        + -
//                                 └── parseFactor      * / % **
//                                       └── parseUnary   ! -x
//                                             └── parsePostfix  i++ i--
//                                                   └── parsePrimary  42 "hi" x


func (p *Parser) parseExpression() Expression {
    return p.parseOr()
}

// "a || b || c"
func (p *Parser) parseOr() Expression {
	left := p.parseAnd()

	for p.current().Type == lexer.OR {
		op := p.current().Literal
		line := p.current().Line
		p.advance()
		right := p.parseAnd()
		left = &InfixExpression{
			exprNode: exprNode{Line: line},
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}

	return left
}

// "a && b && c"
func (p *Parser) parseAnd() Expression {
	left := p.parseEquality()

	for p.current().Type == lexer.AND {
		op := p.current().Literal
		line := p.current().Line
		p.advance()
		right := p.parseEquality()
		left = &InfixExpression{
			exprNode: exprNode{Line: line},
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}
	return left
}

// "a == b"   "a != b"
func (p *Parser) parseEquality() Expression {
	left := p.parseComparison()

	for p.current().Type == lexer.EQ || p.current().Type == lexer.NOT_EQ {
		op := p.current().Literal
		line := p.current().Line
		p.advance()
		right := p.parseComparison()
		left = &InfixExpression{
			exprNode: exprNode{Line: line},
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}
	return left
}

// "age >= 18"   "x < 10"
func (p *Parser) parseComparison() Expression {
	left := p.parseTerm()

	for p.current().Type == lexer.LT || p.current().Type == lexer.GT || p.current().Type == lexer.LTE || p.current().Type == lexer.GTE {
		op := p.current().Literal
		line := p.current().Line
		p.advance()
		right := p.parseTerm()
		left = &InfixExpression{
			exprNode: exprNode{Line: line},
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}
	return left
}

// "a + b - c"
func (p *Parser) parseTerm() Expression {
	left := p.parseFactor()

	for p.current().Type == lexer.PLUS || p.current().Type == lexer.MINUS {
		op := p.current().Literal
		line := p.current().Line
		p.advance()
		right := p.parseFactor()
		left = &InfixExpression{
			exprNode: exprNode{Line: line},
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}
	return left
}

// "a * b / c"   "x ** 2"
func (p *Parser) parseFactor() Expression {
	left := p.parseUnary()

	for p.current().Type == lexer.ASTERISK || p.current().Type == lexer.SLASH || p.current().Type == lexer.PERCENT || p.current().Type == lexer.POWER {
		op := p.current().Literal
		line := p.current().Line
		p.advance()
		right := p.parseUnary()
		left = &InfixExpression{
			exprNode: exprNode{Line: line},
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}
	return left
}

// "!done"   "-x"   "!!x"
func (p *Parser) parseUnary() Expression {
	tok := p.current()

	if tok.Type == lexer.NOT || tok.Type == lexer.MINUS {
		p.advance()
		right := p.parseUnary()
		return &PrefixExpression{
				exprNode: exprNode{Line: tok.Line}, 
				Operator: tok.Literal, 
				Right: right,
		}

	}
	return p.parsePostfix()
}

//i++ and i--
func (p *Parser) parsePostfix() Expression {
	left := p.parsePrimary()

	if p.current().Type == lexer.INCREMENT || p.current().Type == lexer.DECREMENT {
		op := p.current().Literal  
		line := p.current().Line
		p.advance()
		return &PostfixExpression{
			exprNode: exprNode{Line: line},
			Operator: op,
			Left:     left,
		}
	}
	return left
}

func (p *Parser) parsePrimary() Expression {
	tok := p.current()

	switch tok.Type {
	// 42
    case lexer.INT:
		p.advance()
		val, err := strconv.ParseInt(tok.Literal, 10, 64)
		if err != nil {
			p.errors = append(p.errors, fmt.Sprintf(
				"[Skill Issue] \ninvalid number '%s' at line %d", tok.Literal, tok.Line,
			))
			return nil
		}
		return &IntegerLiteral{exprNode: exprNode{Line: tok.Line}, Value: val}
	
	// 3.14
	case lexer.FLOAT:
		p.advance()
		val, err := strconv.ParseFloat(tok.Literal, 64)
		if err != nil {
			p.errors = append(p.errors, fmt.Sprintf(
				"[Skill Issue] \ninvalid number '%s' at line %d", tok.Literal, tok.Line,
			))
			return nil
		}
		return &FloatLiteral{exprNode: exprNode{Line: tok.Line}, Value: val}

	// "hello"
	case lexer.STRING:
		p.advance()
		return &StringLiteral{exprNode: exprNode{Line: tok.Line}, Value: tok.Literal}

	
	// fr_fr → true
    case lexer.TRUE:
        p.advance()
        return &BoolLiteral{exprNode: exprNode{Line: tok.Line}, Value: true}

    // cap → false
    case lexer.FALSE:
        p.advance()
        return &BoolLiteral{exprNode: exprNode{Line: tok.Line}, Value: false}

	// ghosted → nil
    case lexer.NIL:
        p.advance()
        return &NilLiteral{exprNode: exprNode{Line: tok.Line}}

	// x  OR  add(1, 2)
	case lexer.IDENT:
		p.advance()
		if p.current().Type == lexer.LPAREN {
            return p.parseCallExpression(tok) // tok carries the function name
        }
		 if p.current().Type == lexer.LBRACKET {
            ident := &Identifier{exprNode: exprNode{Line: tok.Line}, Name: tok.Literal}
            return p.parseIndexExpression(ident)
        }

		return &Identifier{exprNode: exprNode{Line: tok.Line}, Name: tok.Literal}


	// [1, 2, 3]
    case lexer.LBRACKET:
        return p.parseArrayLiteral()


    // (x + 1)  — grouped expression
    case lexer.LPAREN:
        p.advance()  
        expr := p.parseExpression()
        p.expect(lexer.RPAREN) 
        return expr 

	default:
        p.errors = append(p.errors, fmt.Sprintf(
            "[Skill Issue] unexpected token '%s' at line %d",
            tok.Literal, tok.Line,
        ))
        p.advance()
        return nil
	
	}
}


func (p *Parser) parseCallExpression(fn lexer.Token) Expression {
	call := &CallExpression{
		exprNode: exprNode{Line: fn.Line},
		Function: &Identifier{exprNode: exprNode{Line: fn.Line}, Name: fn.Literal},
	}
 
	p.expect(lexer.LPAREN)
 
	// no arguments: add()
	if p.current().Type == lexer.RPAREN {
		p.advance()
		return call
	}
 
	// first argument
	call.Arguments = append(call.Arguments, p.parseExpression())
 
	// more arguments separated by commas: add(1, 2, 3)
	for p.current().Type == lexer.COMMA {
		p.advance()
		call.Arguments = append(call.Arguments, p.parseExpression())
	}
 
	p.expect(lexer.RPAREN)
	return call
}


func (p *Parser) parseArrayLiteral() Expression {
	arr := &ArrayLiteral{exprNode: exprNode{Line: p.current().Line}}
	p.expect(lexer.LBRACKET)
 
	// empty array: []
	if p.current().Type == lexer.RBRACKET {
		p.advance()
		return arr
	}
 
	// first element
	arr.Elements = append(arr.Elements, p.parseExpression())
 
	// more elements: [1, 2, 3]
	for p.current().Type == lexer.COMMA {
		p.advance()
		arr.Elements = append(arr.Elements, p.parseExpression())
	}
 
	p.expect(lexer.RBRACKET) 

	return arr
}

func (p *Parser) parseIndexExpression(left Expression) Expression {
	idx := &IndexExpression{exprNode: exprNode{Line: p.current().Line}, Left: left}
	p.expect(lexer.LBRACKET)
	idx.Index = p.parseExpression()
	p.expect(lexer.RBRACKET)

	return idx
}