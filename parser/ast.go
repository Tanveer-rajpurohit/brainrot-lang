package parser

//  INTERFACES  — every node must satisfy one

// ANY node that is an expression must implement this
type Expression interface {
	expressionNode()
	GetLine() int
}

// ANY node that is a statement must implement this
type Statement interface {
	statementNode()
	GetLine() int
}

//  ROOT NODE
// Program is the root of every AST — holds all top-level statements
type Program struct {
	Statements []Statement
}

//  STATEMENTS

// trust_me_bro x = 42
type VarStatement struct {
	Line  int
	Name  string
	Value Expression
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) GetLine() int   { return vs.Line }

// x = 42  or  x += 1  (reassignment, not declaration)
type AssignStatement struct {
	Line     int
	Name     string
	Operator string // "=", "+=", "-="
	Value    Expression
}

func (as *AssignStatement) statementNode() {}
func (as *AssignStatement) GetLine() int   { return as.Line }

// say_my_name("hello")
type PrintStatement struct {
	Line  int
	Value Expression
}

func (ps *PrintStatement) statementNode() {}
func (ps *PrintStatement) GetLine() int   { return ps.Line }

// take_this x + 1
type ReturnStatement struct {
	Line  int
	Value Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) GetLine() int   { return rs.Line }

// mission_abort
type BreakStatement struct {
	Line int
}

func (bs *BreakStatement) statementNode() {}
func (bs *BreakStatement) GetLine() int   { return bs.Line }

// skip_this_one
type ContinueStatement struct {
	Line int
}

func (cs *ContinueStatement) statementNode() {}
func (cs *ContinueStatement) GetLine() int   { return cs.Line }

// { ... }  — a block of statements used inside if/while/func
type BlockStatement struct {
	Line       int
	Statements []Statement
}

func (bl *BlockStatement) statementNode() {}
func (bl *BlockStatement) GetLine() int   { return bl.Line }

// chat_is_this_real age >= 18 { ... } nah_bro { ... }
type IfStatement struct {
	Line        int
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement // nil if no nah_bro branch
}

func (is *IfStatement) statementNode() {}
func (is *IfStatement) GetLine() int   { return is.Line }

// on_repeat x < 10 { ... }
type WhileStatement struct {
	Line      int
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode() {}
func (ws *WhileStatement) GetLine() int   { return ws.Line }

// run_it_back i = 0; i < 10; i += 1 { ... }
type ForStatement struct {
	Line      int
	Init      Statement  // trust_me_bro i = 0
	Condition Expression // i < 10
	Post      Statement  // i += 1
	Body      *BlockStatement
}

func (fs *ForStatement) statementNode() {}
func (fs *ForStatement) GetLine() int   { return fs.Line }

// let_him_cook add(a, b) { ... }
type FuncStatement struct {
	Line   int
	Name   string
	Params []string // parameter names
	Body   *BlockStatement
}

func (fn *FuncStatement) statementNode() {}
func (fn *FuncStatement) GetLine() int   { return fn.Line }

// a bare expression used as a statement, e.g. a function call: add(1, 2)
type ExpressionStatement struct {
	Line  int
	Value Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) GetLine() int   { return es.Line }



//  EXPRESSIONS  (things that produce a value)
// 42
type IntegerLiteral struct {
	Line  int
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) GetLine() int    { return il.Line }

// 3.14
type FloatLiteral struct {
	Line  int
	Value float64
}

func (fl *FloatLiteral) expressionNode() {}
func (fl *FloatLiteral) GetLine() int    { return fl.Line }

// "hello"
type StringLiteral struct {
	Line  int
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) GetLine() int    { return sl.Line }

// fr_fr  or  cap
type BoolLiteral struct {
	Line  int
	Value bool
}

func (bl *BoolLiteral) expressionNode() {}
func (bl *BoolLiteral) GetLine() int    { return bl.Line }

// ghosted
type NilLiteral struct {
	Line int
}

func (nl *NilLiteral) expressionNode() {}
func (nl *NilLiteral) GetLine() int    { return nl.Line }

// x  (a variable name used in an expression)
type Identifier struct {
	Line int
	Name string
}

func (id *Identifier) expressionNode() {}
func (id *Identifier) GetLine() int    { return id.Line }

// x + y  |  x == y  |  x && y  etc.
type InfixExpression struct {
	Line     int
	Left     Expression
	Operator string // "+", "-", "==", "!=", "<", ">", "&&", "||", ...
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) GetLine() int    { return ie.Line }

// -x  or  !x
type PrefixExpression struct {
	Line     int
	Operator string // "-" or "!"
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) GetLine() int    { return pe.Line }

// add(1, 2)
type CallExpression struct {
	Line      int
	Function  Expression // usually an Identifier, but can be any expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) GetLine() int    { return ce.Line }

// [1, 2, 3]
type ArrayLiteral struct {
	Line     int
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) GetLine() int    { return al.Line }

// arr[0]
type IndexExpression struct {
	Line  int
	Left  Expression // the array/object
	Index Expression // the index
}

func (ix *IndexExpression) expressionNode() {}
func (ix *IndexExpression) GetLine() int    { return ix.Line }
