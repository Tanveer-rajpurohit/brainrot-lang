package utils

import (
	"fmt"
	"strings"

	"brainrot-lang/lexer"
	"brainrot-lang/parser"
)

// PrintProgram prints the full AST to the terminal in a readable tree format
func PrintProgram(program *parser.Program) {
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

func printStatement(stmt parser.Statement, prefix string, isLast bool) {
	connector, childPrefix := branch(isLast)

	switch s := stmt.(type) {
	case *parser.VarStatement:
		fmt.Printf("%s%sVarStatement (line %d)\n", prefix, connector, s.Line)
		fmt.Printf("%s%s├── Name: %s\n", prefix, childPrefix, s.Name)
		printExprLabel(s.Value, prefix+childPrefix, "Value", true)

	case *parser.AssignStatement:
		fmt.Printf("%s%sAssignStatement (line %d)\n", prefix, connector, s.Line)
		fmt.Printf("%s%s├── Name: %s\n", prefix, childPrefix, s.Name)
		fmt.Printf("%s%s├── Op:   %s\n", prefix, childPrefix, s.Operator)
		printExprLabel(s.Value, prefix+childPrefix, "Value", true)

	case *parser.PrintStatement:
		fmt.Printf("%s%sPrintStatement (line %d)\n", prefix, connector, s.Line)
		printExprLabel(s.Value, prefix+childPrefix, "Value", true)

	case *parser.ReturnStatement:
		fmt.Printf("%s%sReturnStatement (line %d)\n", prefix, connector, s.Line)
		if s.Value != nil {
			printExprLabel(s.Value, prefix+childPrefix, "Value", true)
		}

	case *parser.BreakStatement:
		fmt.Printf("%s%sBreakStatement (line %d)\n", prefix, connector, s.Line)

	case *parser.ContinueStatement:
		fmt.Printf("%s%sContinueStatement (line %d)\n", prefix, connector, s.Line)

	case *parser.ExpressionStatement:
		fmt.Printf("%s%sExpressionStatement (line %d)\n", prefix, connector, s.Line)
		printExprLabel(s.Value, prefix+childPrefix, "Expr", true)

	case *parser.IfStatement:
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

	case *parser.WhileStatement:
		fmt.Printf("%s%sWhileStatement (line %d)\n", prefix, connector, s.Line)
		printExprLabel(s.Condition, prefix+childPrefix, "Condition", false)
		printBlockLabel(s.Body, prefix+childPrefix, "Body", true)

	case *parser.ForStatement:
		fmt.Printf("%s%sForStatement (line %d)\n", prefix, connector, s.Line)
		if s.Init != nil {
			printStatement(s.Init, prefix+childPrefix, false)
		}
		printExprLabel(s.Condition, prefix+childPrefix, "Condition", false)
		if s.Post != nil {
			printStatement(s.Post, prefix+childPrefix, false)
		}
		printBlockLabel(s.Body, prefix+childPrefix, "Body", true)

	case *parser.FuncStatement:
		fmt.Printf("%s%sFuncStatement (line %d)  %s(%s)\n", prefix, connector, s.Line, s.Name, strings.Join(s.Params, ", "))
		printBlockLabel(s.Body, prefix+childPrefix, "Body", true)

	case *parser.BlockStatement:
		fmt.Printf("%s%sBlock (line %d)\n", prefix, connector, s.Line)
		for i, inner := range s.Statements {
			printStatement(inner, prefix+childPrefix, i == len(s.Statements)-1)
		}

	default:
		fmt.Printf("%s%s<unknown statement>\n", prefix, connector)
	}
}

func printElseIf(ei *parser.ElseIfClause, prefix string, isLast bool) {
	connector, childPrefix := branch(isLast)
	fmt.Printf("%s%sElseIf (line %d)\n", prefix, connector, ei.Line)
	printExprLabel(ei.Condition, prefix+childPrefix, "Condition", false)
	printBlockLabel(ei.Body, prefix+childPrefix, "Body", true)
}

func printBlockLabel(block *parser.BlockStatement, prefix, label string, isLast bool) {
	if block == nil {
		return
	}
	connector, childPrefix := branch(isLast)
	fmt.Printf("%s%s%s: Block\n", prefix, connector, label)
	for i, s := range block.Statements {
		printStatement(s, prefix+childPrefix, i == len(block.Statements)-1)
	}
}

func printExprLabel(expr parser.Expression, prefix, label string, isLast bool) {
	connector, childPrefix := branch(isLast)
	fmt.Printf("%s%s%s: %s\n", prefix, connector, label, exprString(expr, prefix+childPrefix))
}

func exprString(expr parser.Expression, _ string) string {
	if expr == nil {
		return "<nil>"
	}
	switch e := expr.(type) {
	case *parser.IntegerLiteral:
		return fmt.Sprintf("Int(%d)", e.Value)
	case *parser.FloatLiteral:
		return fmt.Sprintf("Float(%g)", e.Value)
	case *parser.StringLiteral:
		return fmt.Sprintf("String(%q)", e.Value)
	case *parser.BoolLiteral:
		return fmt.Sprintf("Bool(%v)", e.Value)
	case *parser.NilLiteral:
		return "Nil"
	case *parser.Identifier:
		return fmt.Sprintf("Ident(%s)", e.Name)
	case *parser.InfixExpression:
		return fmt.Sprintf("(%s %s %s)", exprString(e.Left, ""), e.Operator, exprString(e.Right, ""))
	case *parser.PrefixExpression:
		return fmt.Sprintf("(%s%s)", e.Operator, exprString(e.Right, ""))
	case *parser.PostfixExpression:
		return fmt.Sprintf("(%s%s)", exprString(e.Left, ""), e.Operator)
	case *parser.CallExpression:
		args := make([]string, len(e.Arguments))
		for i, a := range e.Arguments {
			args[i] = exprString(a, "")
		}
		return fmt.Sprintf("Call(%s, [%s])", exprString(e.Function, ""), strings.Join(args, ", "))
	case *parser.ArrayLiteral:
		elems := make([]string, len(e.Elements))
		for i, el := range e.Elements {
			elems[i] = exprString(el, "")
		}
		return fmt.Sprintf("Array[%s]", strings.Join(elems, ", "))
	case *parser.IndexExpression:
		return fmt.Sprintf("%s[%s]", exprString(e.Left, ""), exprString(e.Index, ""))
	default:
		return "<expr>"
	}
}

// PrintLexicalTable prints a formatted lexical token table
func PrintLexicalTable(tokens []lexer.Token) {
	// Print header
	fmt.Printf("\n%s┌─────┬──────────────┬──────────────┬─────────────────┬────────┐%s\n", ColorBlue, ColorReset)
	fmt.Printf("%s│ IDX │ TOKEN TYPE   │ LITERAL      │ CATEGORY        │ POS    │%s\n", ColorBlue, ColorReset)
	fmt.Printf("%s├─────┼──────────────┼──────────────┼─────────────────┼────────┤%s\n", ColorBlue, ColorReset)

	// Print each token row
	for i, tok := range tokens {
		if tok.Type == lexer.NEWLINE {
			continue // skip newlines
		}

		// Get token category
		category := GetTokenCategory(tok.Type)

		// Format index
		idx := fmt.Sprintf("[%d]", i)

		// Format token type
		tokType := fmt.Sprintf("%s", tok.Type)

		// Format literal (truncate if too long)
		literal := tok.Literal
		if len(literal) > 12 {
			literal = literal[:9] + "..."
		}
		if tok.Type == lexer.EOF {
			literal = "EOF"
		}

		// Format position
		pos := fmt.Sprintf("L%d:C%d", tok.Line, tok.Column)

		// Print row with colors
		fmt.Printf("%s│ %-3s │ %s%-12s%s │ %s%-12s%s │ %s%-15s%s │ %-6s │%s\n",
			ColorBlue,
			idx,
			ColorCyan, tokType, ColorReset,
			ColorGreen, literal, ColorReset,
			ColorYellow, category, ColorReset,
			pos,
			ColorBlue)
	}

	// Print footer
	fmt.Printf("%s└─────┴──────────────┴──────────────┴─────────────────┴────────┘%s\n", ColorBlue, ColorReset)
	fmt.Printf("\n%sTotal Tokens: %d%s\n\n", ColorBold, len(tokens), ColorReset)
}

func PrintHelp() {
	fmt.Println(`
Usage: brainrot <command> <file.brt>
 
Commands:
  run     <file.brt>   Run a BrainRot program
  tokens  <file.brt>   Debug: show token stream only
  ast     <file.brt>   Debug: show AST tree only
  help                 Show this message
 
Examples:
  brainrot run examples/hello.brt
  brainrot tokens examples/hello.brt
  brainrot ast examples/hello.brt
 
Keywords:
  trust_me_bro        → variable
  say_my_name(...)    → print
  chat_is_this_real   → if
  wait_hold_up        → else if
  nah_bro             → else
  on_repeat           → while
  run_it_back         → for
  let_him_cook        → function
  take_this           → return
  mission_abort       → break
  skip_this_one       → continue
  fr_fr / cap         → true / false
  ghosted             → nil
	`)
}
