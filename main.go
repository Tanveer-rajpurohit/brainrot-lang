package main

import (
	"fmt"
	"os"

	"brainrot-lang/interpreter"
	"brainrot-lang/lexer"
	"brainrot-lang/parser"
	"brainrot-lang/utils"
)

func main() {
	fmt.Println("BrainRot Lang — Compiler Booting... 🔥")

	utils.Banner()

	// cml run command go run main.go run examples/hello.brl in no prod ./brainrot run examples/hello.brl
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [arguments]")
		os.Exit(1)
	}

	command := os.Args[1]

	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Println(os.Args[i])
	// }

	switch command {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Usage: brainrot run <file.brl>")
			os.Exit(1)
		}
		runFile(os.Args[2], false, false)
 
	case "tokens":
		if len(os.Args) < 3 {
			fmt.Println("Usage: brainrot tokens <file.brl>")
			os.Exit(1)
		}
		runFile(os.Args[2], true, false)
 
	case "ast":
		if len(os.Args) < 3 {
			fmt.Println("Usage: brainrot ast <file.brl>")
			os.Exit(1)
		}
		runFile(os.Args[2], false, true)
 
	case "help":
		utils.PrintHelp()

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func runFile(filename string, showTokens bool, showAST bool) {
	source, err := os.ReadFile(filename)
	if err != nil {
		utils.Fatal(utils.NewError("CLI", 0, 0, fmt.Sprintf("system open file '%s'", filename)))
		os.Exit(1)
	}

	
	// Step 1: Lexer
	l := lexer.New(string(source))
	tokens := l.Tokenize()
 
	if showTokens {
		fmt.Printf("%s[LEXER OUTPUT]%s\n", utils.ColorGreen, utils.ColorReset)
		utils.PrintLexicalTable(tokens)
	}

	// Step 2: Parser
	p := parser.New(tokens)
	program := p.Parse()
 
	if errs := p.Errors(); len(errs) > 0 {
		fmt.Printf("%s[PARSER ERRORS]%s\n", utils.ColorRed, utils.ColorReset)
		for _, e := range errs {
			fmt.Println(e)
		}
		os.Exit(1)
	}
 
	if showAST {
		fmt.Printf("%s[AST OUTPUT]%s\n", utils.ColorGreen, utils.ColorReset)
		utils.PrintProgram(program)
		return
	}
 
	// Step 3: Interpreter
	interp := interpreter.New()
	interp.Eval(program)
 
	if errs := interp.Errors(); len(errs) > 0 {
		fmt.Printf("\n%s%s\n", utils.ColorRed, utils.ColorReset)
		for _, e := range errs {
			fmt.Printf("%s%s%s\n", utils.ColorRed, e, utils.ColorReset)
		}
		os.Exit(1)
	}
}



