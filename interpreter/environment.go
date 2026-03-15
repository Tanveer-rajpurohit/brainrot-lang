package interpreter

import "fmt"

// Environment is the symbol table — stores variables for one scope level
type Environment struct {
    store map[string]interface{} // variable name → value
    outer *Environment           // parent scope (nil for global)
}


func NewEnvironment() *Environment {
    return &Environment{
        store: make(map[string]interface{}),
        outer: nil,
    }
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env

}

func (e *Environment) Get(name string) (interface{}, bool) {
    val, ok := e.store[name]
    if !ok && e.outer != nil {
        val, ok = e.outer.Get(name)
    }
    return val, ok
}

func (e *Environment) Set(name string, val interface{}) {
    e.store[name] = val
}

func (e *Environment) Update(name string, val interface{}) bool {
    if _, ok := e.store[name]; ok {
        e.store[name] = val
        return true
    }
    if e.outer != nil {
        return e.outer.Update(name, val) 
    }
    return false
}

func (e *Environment) String() string {
    result := "Environment{\n"
    for k, v := range e.store {
        result += fmt.Sprintf("  %s = %v\n", k, v)
    }
    if e.outer != nil {
        result += "  outer: " + e.outer.String()
    }
    result += "}"
    return result
}