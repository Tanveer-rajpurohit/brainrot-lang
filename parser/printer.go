package parser

import (
	"fmt"
	"strings"
)

// PrintProgram prints the full AST to the terminal in a readable tree format
func PrintProgram(program *Program) {
	fmt.Println("Program")
	for i, stmt := range program.Statements {
		isLast := i == len(program.Statements)-1
		printStatement(stmt, "", isLast)
	}
}

func branch(isLast bool) (connector, childPrefix string) {
	if isLast {
		return "└── ", "    "
	}
	return "├── ", "│   "
}

func printStatement(stmt Statement, prefix string, isLast bool) {
	connector, childPrefix := branch(isLast)

	switch s := stmt.(type) {
	case *VarStatement:
		fmt.Printf("%s%sVarStatement (line %d)\n", prefix, connector, s.Line)
		fmt.Printf("%s%s├── Name: %s\n", prefix, childPrefix, s.Name)
		printExprLabel(s.Value, prefix+childPrefix, "Value", true)

	case *AssignStatement:
		fmt.Printf("%s%sAssignStatement (line %d)\n", prefix, connector, s.Line)
		fmt.Printf("%s%s├── Name: %s\n", prefix, childPrefix, s.Name)
		fmt.Printf("%s%s├── Op:   %s\n", prefix, childPrefix, s.Operator)
		printExprLabel(s.Value, prefix+childPrefix, "Value", true)

	case *PrintStatement:
		fmt.Printf("%s%sPrintStatement (line %d)\n", prefix, connector, s.Line)
		printExprLabel(s.Value, prefix+childPrefix, "Value", true)

	case *ReturnStatement:
		fmt.Printf("%s%sReturnStatement (line %d)\n", prefix, connector, s.Line)
		if s.Value != nil {
			printExprLabel(s.Value, prefix+childPrefix, "Value", true)
		}

	case *BreakStatement:
		fmt.Printf("%s%sBreakStatement (line %d)\n", prefix, connector, s.Line)

	case *ContinueStatement:
		fmt.Printf("%s%sContinueStatement (line %d)\n", prefix, connector, s.Line)

	case *ExpressionStatement:
		fmt.Printf("%s%sExpressionStatement (line %d)\n", prefix, connector, s.Line)
		printExprLabel(s.Value, prefix+childPrefix, "Expr", true)

	case *IfStatement:
		fmt.Printf("%s%sIfStatement (line %d)\n", prefix, connector, s.Line)
		printExprLabel(s.Condition, prefix+childPrefix, "Condition", false)
		printBlockLabel(s.Consequence, prefix+childPrefix, "Then", len(s.ElseIf) == 0 && s.Alternative == nil)
		for i, ei := range s.ElseIf {
			last := i == len(s.ElseIf)-1 && s.Alternative == nil
			printElseIf(ei, prefix+childPrefix, last)
		}
		if s.Alternative != nil {
			printBlockLabel(s.Alternative, prefix+childPrefix, "Else", true)
		}

	case *WhileStatement:
		fmt.Printf("%s%sWhileStatement (line %d)\n", prefix, connector, s.Line)
		printExprLabel(s.Condition, prefix+childPrefix, "Condition", false)
		printBlockLabel(s.Body, prefix+childPrefix, "Body", true)

	case *ForStatement:
		fmt.Printf("%s%sForStatement (line %d)\n", prefix, connector, s.Line)
		if s.Init != nil {
			printStatement(s.Init, prefix+childPrefix, false)
		}
		printExprLabel(s.Condition, prefix+childPrefix, "Condition", false)
		if s.Post != nil {
			printStatement(s.Post, prefix+childPrefix, false)
		}
		printBlockLabel(s.Body, prefix+childPrefix, "Body", true)

	case *FuncStatement:
		fmt.Printf("%s%sFuncStatement (line %d)  %s(%s)\n", prefix, connector, s.Line, s.Name, strings.Join(s.Params, ", "))
		printBlockLabel(s.Body, prefix+childPrefix, "Body", true)

	case *BlockStatement:
		fmt.Printf("%s%sBlock (line %d)\n", prefix, connector, s.Line)
		for i, inner := range s.Statements {
			printStatement(inner, prefix+childPrefix, i == len(s.Statements)-1)
		}

	default:
		fmt.Printf("%s%s<unknown statement>\n", prefix, connector)
	}
}

func printElseIf(ei *ElseIfClause, prefix string, isLast bool) {
	connector, childPrefix := branch(isLast)
	fmt.Printf("%s%sElseIf (line %d)\n", prefix, connector, ei.Line)
	printExprLabel(ei.Condition, prefix+childPrefix, "Condition", false)
	printBlockLabel(ei.Body, prefix+childPrefix, "Body", true)
}

func printBlockLabel(block *BlockStatement, prefix, label string, isLast bool) {
	if block == nil {
		return
	}
	connector, childPrefix := branch(isLast)
	fmt.Printf("%s%s%s: Block\n", prefix, connector, label)
	for i, s := range block.Statements {
		printStatement(s, prefix+childPrefix, i == len(block.Statements)-1)
	}
}

func printExprLabel(expr Expression, prefix, label string, isLast bool) {
	connector, childPrefix := branch(isLast)
	fmt.Printf("%s%s%s: %s\n", prefix, connector, label, exprString(expr, prefix+childPrefix))
}

func exprString(expr Expression, _ string) string {
	if expr == nil {
		return "<nil>"
	}
	switch e := expr.(type) {
	case *IntegerLiteral:
		return fmt.Sprintf("Int(%d)", e.Value)
	case *FloatLiteral:
		return fmt.Sprintf("Float(%g)", e.Value)
	case *StringLiteral:
		return fmt.Sprintf("String(%q)", e.Value)
	case *BoolLiteral:
		return fmt.Sprintf("Bool(%v)", e.Value)
	case *NilLiteral:
		return "Nil"
	case *Identifier:
		return fmt.Sprintf("Ident(%s)", e.Name)
	case *InfixExpression:
		return fmt.Sprintf("(%s %s %s)", exprString(e.Left, ""), e.Operator, exprString(e.Right, ""))
	case *PrefixExpression:
		return fmt.Sprintf("(%s%s)", e.Operator, exprString(e.Right, ""))
	case *PostfixExpression:
		return fmt.Sprintf("(%s%s)", exprString(e.Left, ""), e.Operator)
	case *CallExpression:
		args := make([]string, len(e.Arguments))
		for i, a := range e.Arguments {
			args[i] = exprString(a, "")
		}
		return fmt.Sprintf("Call(%s, [%s])", exprString(e.Function, ""), strings.Join(args, ", "))
	case *ArrayLiteral:
		elems := make([]string, len(e.Elements))
		for i, el := range e.Elements {
			elems[i] = exprString(el, "")
		}
		return fmt.Sprintf("Array[%s]", strings.Join(elems, ", "))
	case *IndexExpression:
		return fmt.Sprintf("%s[%s]", exprString(e.Left, ""), exprString(e.Index, ""))
	default:
		return "<expr>"
	}
}
