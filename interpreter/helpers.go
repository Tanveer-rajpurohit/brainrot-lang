package interpreter

import (
	"fmt"
	"math"
)

// isTruthy decides if a value counts as "true" in a condition
func isTruthy(val interface{}) bool {
	switch v := val.(type) {
		case bool:
			return v
		case int64:
			return v != 0
		case float64:
			return v != 0.0
		case string:
			return v != ""
		case nil:
			return false
		default:
			return true
	}
}

// formatValue converts any runtime value to a printable string
func formatValue(val interface{}) string {
	if val == nil {
		return "ghosted"
	}
	switch v := val.(type) {
	case bool:
		if v {
			return "fr fr"
		}
		return "cap"
	case int64:
		return fmt.Sprintf("%d", v)
	case float64:
		if v == math.Trunc(v) {
			return fmt.Sprintf("%.0f", v)
		}
		return fmt.Sprintf("%g", v)
	case string:
		return v
	case *FuncValue:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

// evalInfix performs the actual math / comparison between two values
func evalInfix(left interface{}, op string, right interface{}) interface{} {
	if l, lok := left.(int64); lok {
		if r, rok := right.(int64); rok {
			switch op {
			case "+":
				return l + r
			case "-":
				return l - r
			case "*":
				return l * r
			case "/":
				if r == 0 {
					return nil
				}
				return l / r
			case "%":
				return l % r
			case "**":
				return int64(math.Pow(float64(l), float64(r)))
			case "==":
				return l == r
			case "!=":
				return l != r
			case "<":
				return l < r
			case ">":
				return l > r
			case "<=":
				return l <= r
			case ">=":
				return l >= r
			}
		}

		if r, rok := right.(float64); rok {
			lf := float64(l)
			switch op {
			case "+":
				return lf + r
			case "-":
				return lf - r
			case "*":
				return lf * r
			case "/":
				if r == 0 {
					return nil
				}
				return lf / r
			case "**":
				return math.Pow(lf, r)
			case "==":
				return lf == r
			case "!=":
				return lf != r
			case "<":
				return lf < r
			case ">":
				return lf > r
			case "<=":
				return lf <= r
			case ">=":
				return lf >= r
			}
		}
	}

	if l, lok := left.(float64); lok {
		var r float64
		switch rv := right.(type) {
		case float64:
			r = rv
		case int64:
			r = float64(rv)
		default:
			return nil
		}
		switch op {
		case "+":
			return l + r
		case "-":
			return l - r
		case "*":
			return l * r
		case "/":
			if r == 0 {
				return nil
			}
			return l / r
		case "**":
			return math.Pow(l, r)
		case "==":
			return l == r
		case "!=":
			return l != r
		case "<":
			return l < r
		case ">":
			return l > r
		case "<=":
			return l <= r
		case ">=":
			return l >= r
		}
	}

	if l, lok := left.(bool); lok {
		if r, rok := right.(bool); rok {
			switch op {
			case "==":
				return l == r
			case "!=":
				return l != r
			case "&&":
				return l && r
			case "||":
				return l || r
			}
		}
	}

	if l, lok := left.(string); lok {
		switch op {
		case "+":
			return l + formatValue(right)
		case "==":
			if r, rok := right.(string); rok {
				return l == r
			}
		case "!=":
			if r, rok := right.(string); rok {
				return l != r
			}
		}
	}

	switch op {
	case "&&":
		return isTruthy(left) && isTruthy(right)
	case "||":
		return isTruthy(left) || isTruthy(right)
	case "==":
		return left == right
	case "!=":
		return left != right
	}

	return nil
}