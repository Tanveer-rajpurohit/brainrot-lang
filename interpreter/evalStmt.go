package interpreter

import (
	"brainrot-lang/parser"
	"fmt"
)

// Statements DO things: declare a var, print, run an if, call a loop
func (i *Interpreter) evalStatement(node parser.Statement) interface{} {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *parser.ExpressionStatement:
		return i.EvalExpr(n.Value)
	
	case *parser.VarStatement:
		val := i.EvalExpr(n.Value)
		i.env.Set(n.Name, val)
		return nil

	case *parser.AssignStatement:
		val := i.EvalExpr(n.Value)
		switch n.Operator{

		case "=":
			if !i.env.Update(n.Name, val) {
				i.env.Set(n.Name, val)
			}
		case "+=":
			old, _ := i.env.Get(n.Name)
			i.env.Update(n.Name, evalInfix(old, "+", val))
		case "-=":
			old, _ := i.env.Get(n.Name)
			i.env.Update(n.Name, evalInfix(old, "-", val))
		case "*=":
			old, _ := i.env.Get(n.Name)
			i.env.Update(n.Name, evalInfix(old, "*", val))
		case "/=":
			old, _ := i.env.Get(n.Name)
			i.env.Update(n.Name, evalInfix(old, "/", val))
		
		}
		return nil

	case *parser.PrintStatement:
		val := i.EvalExpr(n.Value)
		fmt.Println(formatValue(val))
		return nil
	
	case *parser.BlockStatement:
		return i.evalBlock(n)
	
	case *parser.IfStatement:
		return i.evalIf(n)

	case *parser.WhileStatement:
		return i.evalWhile(n)

	case *parser.ForStatement:
		return i.evalFor(n)

	case *parser.FuncStatement:
		fn := &FuncValue{
			Params: n.Params,
			Body:   n.Body,
			Env:    i.env,
		}
		i.env.Set(n.Name, fn)
		return nil

	case *parser.ReturnStatement:
		if n.Value != nil {
			return &ReturnValue{Value: i.EvalExpr(n.Value)}
		}
		return &ReturnValue{Value: nil}

	case *parser.BreakStatement:
		return &BreakSignal{}

	case *parser.ContinueStatement:
		return &ContinueSignal{}
	}

	return nil
}


// It first registers all functions, then either calls main() or runs top-level stmts
func (i *Interpreter) evalProgram(program *parser.Program) interface{} {
	// Pass 1: validate + detect duplicates
	seenVars  := make(map[string]int) // varName  → line it was declared
	seenFuncs := make(map[string]int) // funcName → line it was defined

	for _, stmt := range program.Statements {
		switch s := stmt.(type) {

		case *parser.VarStatement:
			if prevLine, exists := seenVars[s.Name]; exists {
				i.runtimeError(s.GetLine(), fmt.Sprintf(
					"global variable '%s' already declared at line %d — no cap, you can't declare it twice",
					s.Name, prevLine,
				))
			} else {
				seenVars[s.Name] = s.GetLine()
			}

		case *parser.FuncStatement:
			if prevLine, exists := seenFuncs[s.Name]; exists {
				i.runtimeError(s.GetLine(), fmt.Sprintf(
					"function '%s' already defined at line %d — let him cook only once bro",
					s.Name, prevLine,
				))
			} else {
				seenFuncs[s.Name] = s.GetLine()
			}

		default:
			i.runtimeError(stmt.GetLine(),
				"not allowed outside main() — only 'trust_me_bro' variables and 'let_him_cook' functions are allowed at top level",
			)
		}
	}

	if len(i.errors) > 0 {
		return nil
	}

	// Pass 2: run global variables top-to-bottom
	for _, stmt := range program.Statements {
		if varStmt, ok := stmt.(*parser.VarStatement); ok {
			val := i.EvalExpr(varStmt.Value)
			i.env.Set(varStmt.Name, val)
		}
	}

	// Pass 3: register all functions (bodies not executed yet)
	for _, stmt := range program.Statements {
		if fnStmt, ok := stmt.(*parser.FuncStatement); ok {
			fn := &FuncValue{
				Params: fnStmt.Params,
				Body:   fnStmt.Body,
				Env:    i.env,
			}
			i.env.Set(fnStmt.Name, fn)
		}
	}

	// Pass 4: call main()
	mainFn, ok := i.env.Get("main")
	if !ok {
		i.runtimeError(0, "no main() found — BrainRot Lang needs 'let_him_cook main() { ... }'")
		return nil
	}

	fn, ok := mainFn.(*FuncValue)
	if !ok {
		i.runtimeError(0, "main is not a function")
		return nil
	}

	mainEnv := NewEnclosedEnvironment(fn.Env)
	oldEnv := i.env
	i.env = mainEnv
	result := i.evalBlock(fn.Body)
	i.env = oldEnv

	if ret, ok := result.(*ReturnValue); ok {
		return ret.Value
	}

	return result
}


// It must stop early and bubble up ReturnValue / BreakSignal / ContinueSignal
func (i *Interpreter) evalBlock(block *parser.BlockStatement) interface{} {
	var result interface{}
	for _, stmt := range block.Statements {
		result = i.evalStatement(stmt)
		switch result.(type) {
		case *ReturnValue, *BreakSignal, *ContinueSignal:
			return result
		}
	}
	return result

}


func (i *Interpreter) evalIf(node *parser.IfStatement) interface{} {
	if isTruthy(i.EvalExpr(node.Condition)) {
		return i.evalBlock(node.Consequence)
	}

	for _, elseIf := range node.ElseIf {
		if isTruthy(i.EvalExpr(elseIf.Condition)) {
			return i.evalBlock(elseIf.Body)
		}
	}

	if node.Alternative != nil {
		return i.evalBlock(node.Alternative)
	}

	return nil

}


func (i *Interpreter) evalWhile(node *parser.WhileStatement) interface{} {
	for isTruthy(i.EvalExpr(node.Condition)) {
		result := i.evalBlock(node.Body)
		switch result.(type) {
		case *BreakSignal:
			return nil
		case *ReturnValue:
			return result
		}
	}
	return nil
}

func (i *Interpreter) evalFor(node *parser.ForStatement) interface{} {

	loopEnv := NewEnclosedEnvironment(i.env)
	oldEnv := i.env
	i.env = loopEnv
	defer func() { i.env = oldEnv }()

	if node.Init != nil {
		i.evalStatement(node.Init)
	}

	for isTruthy(i.EvalExpr(node.Condition)) {
		result := i.evalBlock(node.Body)
		switch result.(type) {
		case *BreakSignal:
			return nil
		case *ReturnValue:
			return result

		}

		if node.Post != nil {
			i.evalStatement(node.Post)
		}
	}
	return nil
}

func (i *Interpreter) evalCall(node *parser.CallExpression) interface{} {

	fnVal := i.EvalExpr(node.Function)
	fn, ok := fnVal.(*FuncValue)
	if !ok {
		i.runtimeError(node.GetLine(), fmt.Sprintf(
			"'%v' is not a function bro", fnVal,
		))
		return nil
	}

	args := make([]interface{}, len(node.Arguments))
	for idx, arg := range node.Arguments {
		args[idx] = i.EvalExpr(arg)
	}

	if len(args) != len(fn.Params) {
		i.runtimeError(node.GetLine(), fmt.Sprintf(
			"function expects %d args but got %d",
			len(fn.Params), len(args),
		))
		return nil
	}

	funcEnv := NewEnclosedEnvironment(fn.Env)
	for idx, param := range fn.Params {
		funcEnv.Set(param, args[idx])
	}

	oldEnv := i.env
	i.env = funcEnv
	result := i.evalBlock(fn.Body)
	i.env = oldEnv

	if ret, ok := result.(*ReturnValue); ok {
		return ret.Value
	}
	return nil
}
