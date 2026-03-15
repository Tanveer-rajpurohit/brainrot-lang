package interpreter

import (
	"fmt"
	"strings"

	"brainrot-lang/parser"
)


type ReturnValue struct{ Value interface{} }  // take_this x
type BreakSignal struct{}                      // mission_abort
type ContinueSignal struct{}                   // skip_this_one


type FuncValue struct {
	Params []string
	Body   *parser.BlockStatement
	Env    *Environment 
}

func (f *FuncValue) String() string {
	return fmt.Sprintf("<function(%s)>", strings.Join(f.Params, ", "))
}


// Interpreter

type Interpreter struct {
	env    *Environment
	errors []string
}

func New() *Interpreter {
	return &Interpreter{
		env:    NewEnvironment(),
		errors: []string{},
	}
}

func (i *Interpreter) Errors() []string {
	return i.errors
}

func (i *Interpreter) runtimeError(line int, msg string) {
	i.errors = append(i.errors, fmt.Sprintf(
		"\n[SKILL ISSUE]\n   [Runtime Error]\n  Line %d → %s\n",
		line, msg,
	))
}


func (i *Interpreter) Eval(program *parser.Program) interface{} {
	return i.evalProgram(program)
}
