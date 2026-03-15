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
	for _, stmt := range program.Statements {
		if fn, ok := stmt.(*parser.FuncStatement); ok {
			i.evalStatement(fn)
		}
	}

	mainFn, ok := i.env.Get("main")

	if !ok {

		for _, stmt := range program.Statements {
			result := i.evalStatement(stmt)
			if ret, ok := result.(*ReturnValue); ok {
				return ret.Value
			}
		}
		return nil
	}

	fn, ok := mainFn.(*FuncValue)
	if !ok {
		i.runtimeError(0, "main is not a function")
		return nil
	}
	funcEnv := NewEnclosedEnvironment(fn.Env)
	oldEnv := i.env
	i.env = funcEnv
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
