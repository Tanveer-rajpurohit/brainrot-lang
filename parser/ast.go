package parser

//BASE TYPES  — written ONCE, embedded everywhere
type stmtNode struct { Line int }
func (s stmtNode) statementNode() {}
func (s stmtNode) GetLine() int { return s.Line }

type exprNode  struct { Line int }
func (e exprNode) expressionNode() {}
func (e exprNode) GetLine() int { return e.Line }


//  INTERFACES
type Expression interface {
	expressionNode()
	GetLine() int
}
 
type Statement interface {
	statementNode()
	GetLine() int
}


//ROOT NODE

type Program struct {
	Statements []Statement
}


// STATEMENT NODES

// trust_me_bro x = 42
type VarStatement struct {
	stmtNode  // ← Line + statementNode() + GetLine()
	name  string // x
	value Expression // 42
}

// x = 42   x += 1   x -= 1   (re-assignment, not declaration)
type AssignStatement struct {
	stmtNode
	Name     string
	Operator string // "=" | "+=" | "-="
	Value    Expression
}

// say_my_name("hello world")
type PrintStatement struct {
	stmtNode
	Value Expression
}

// take_this x + 1
type ReturnStatement struct {
	stmtNode
	Value Expression // nil = bare "take_this" with no value
}

// mission_abort
type BreakStatement struct {
	stmtNode
}

// skip_this_one
type ContinueStatement struct {
	stmtNode
}

// { statement1; statement2; ... }
// Used as the body of if / while / for / func
type BlockStatement struct {
	stmtNode
	Statements []Statement
}

// chat_is_this_real age >= 18 { ... }
// wait_hold_up age >= 5  { ... }   ← zero or more
// nah_bro                { ... }   ← optional
type IfStatement struct {
	stmtNode
	Condition   Expression
	Consequence *BlockStatement
	ElseIf      []*ElseIfClause  // wait_hold_up chains (can be empty)
	Alternative *BlockStatement  // nah_bro block (can be nil)
}

// wait_hold_up x > 5 { ... }
type ElseIfClause struct {
	stmtNode
	Condition Expression
	Body      *BlockStatement
}

// on_repeat x < 10 { ... }
type WhileStatement struct {
	stmtNode
	Condition Expression
	Body      *BlockStatement
}

// run_it_back (i = 0; i < 10; i += 1) { ... }
type ForStatement struct {
	stmtNode
	Init      Statement  // trust_me_bro i = 0  OR  i = 0
	Condition Expression // i < 10
	Post      Statement  // i += 1
	Body      *BlockStatement
}
 
// let_him_cook add(a, b) { ... }
type FuncStatement struct {
	stmtNode
	Name   string
	Params []string
	Body   *BlockStatement
}

// add(1, 2)  used as a standalone statement on its own line
// Without this wrapper, CallExpression can't go into []Statement
type ExpressionStatement struct {
	stmtNode
	Value Expression
}


//  EXPRESSIONS NODES


// 42
type IntegerLiteral struct {
	exprNode
	Value int64
}
 
// 3.14
type FloatLiteral struct {
	exprNode
	Value float64
}
 
// "hello walter"
type StringLiteral struct {
	exprNode
	Value string
}
 
// fr_fr  or  cap
type BoolLiteral struct {
	exprNode
	Value bool
}
 
// ghosted
type NilLiteral struct {
	exprNode
}
 
// x  myVar  name  (variable reference in an expression)
type Identifier struct {
	exprNode
	Name string
}

// left OP right   ->   x + y   age >= 18   a == b   x && y
type InfixExpression struct {
	exprNode
	Left     Expression
	Operator string
	Right    Expression
}
 
// OP right   ->   !done   -x
type PrefixExpression struct {
	exprNode
	Operator string
	Right    Expression
}

// OP left  ->   i++   i--
type PostfixExpression struct {
	exprNode
	Operator string
	Left     Expression
}

// callee(arg1, arg2)
type CallExpression struct {
	exprNode
	Function  Expression
	Arguments []Expression
}
 
// [1, 2, 3]
type ArrayLiteral struct {
	exprNode
	Elements []Expression
}
 
// arr[0]   arr[i], arr[i + 1]
type IndexExpression struct {
	exprNode
	Left  Expression // the array -> arr
	Index Expression // the index value -> value inside []
}