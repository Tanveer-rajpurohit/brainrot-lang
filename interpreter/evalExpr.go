package interpreter

import (
	"brainrot-lang/parser"
	"fmt"
)

// 42, x, x + 1, add(1, 2)
func (i *Interpreter) EvalExpr(node parser.Expression) interface{} {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *parser.IntegerLiteral:
		return n.Value
	case *parser.FloatLiteral:
		return n.Value

	case *parser.StringLiteral:
		return n.Value

	case *parser.BoolLiteral:
		return n.Value

	case *parser.NilLiteral:
		return nil

	// x → env.Get("x") → 42
	case *parser.Identifier:
		val, ok := i.env.Get(n.Name)
		if !ok {
			i.runtimeError(n.GetLine(), fmt.Sprintf("variable '%s' is ghosted (not defined)", n.Name))
			return nil
		}
		return val

	// x + 1, age >= 18, a == b
	case *parser.InfixExpression:
		left := i.EvalExpr(n.Left)
		right := i.EvalExpr(n.Right)
		return evalInfix(left, n.Operator, right)

	// !done, -x
	case *parser.PrefixExpression:
		right := i.EvalExpr(n.Right)
		switch n.Operator {
		case "!":
			return !isTruthy(right)
		case "-":
			switch v := right.(type) {
			case int64:
				return -v
			case float64:
				return -v
			}
		}
	// x++ , x--
	case *parser.PostfixExpression:
		ident, ok := n.Left.(*parser.Identifier)
		if !ok {
			i.runtimeError(n.GetLine(), "++ / -- can only be used on variables")
			return nil
		}
		val, exists := i.env.Get(ident.Name)
		if !exists {
			i.runtimeError(n.GetLine(), fmt.Sprintf("variable '%s' is ghosted", ident.Name))
			return nil
		}
		intVal, ok := val.(int64)
		if !ok {
			i.runtimeError(n.GetLine(), "++ / -- can only be used on integers")
			return nil
		}
		if n.Operator == "++" {
			i.env.Update(ident.Name, intVal+1)
		} else {
			i.env.Update(ident.Name, intVal-1)
		}
		return intVal

	// add(1, 2), greet("Walter")
	case *parser.CallExpression:
		return i.evalCall(n)

	// [1, 2, 3], ["a", "b"]
	case *parser.ArrayLiteral:
		elements := make([]interface{}, len(n.Elements))
		for idx, el := range n.Elements {
			elements[idx] = i.EvalExpr(el)
		}
		return elements

	// arr[0], arr[i]
	case *parser.IndexExpression:
		left := i.EvalExpr(n.Left)
		index := i.EvalExpr(n.Index)
		arr, ok := left.([]interface{})
		if !ok {
			i.runtimeError(n.GetLine(), "index operator used on non-array value")
			return nil
		}
		idx, ok := index.(int64)
		if !ok {
			i.runtimeError(n.GetLine(), "array index must be an integer")
			return nil
		}
		if idx < 0 || int(idx) >= len(arr) {
			i.runtimeError(n.GetLine(), fmt.Sprintf("index %d out of bounds (array length %d)", idx, len(arr)))
			return nil
		}
		return arr[idx]
	}

	return nil
}
