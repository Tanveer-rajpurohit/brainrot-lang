package main

import (
	"fmt"
	"os"

	"brainrot-lang/lexer"
	"brainrot-lang/utils"
	"brainrot-lang/parser"
)

func main() {
	fmt.Println("BrainRot Lang — Compiler Booting... 🔥")

	utils.Banner()

	// cml run command go run main.go run examples/hello.brl in no prod ./brainrot run examples/hello.brl
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [arguments]")
	}

	command := os.Args[1]

	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Println(os.Args[i])
	// }

	switch command {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go run <file.brl>")
			os.Exit(1)
		}
		filename := os.Args[2]
		fmt.Printf("Running file: %s\n", filename)
		runFile(filename)

	case "build":

	case "help":

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func runFile(filename string) {
	source, err := os.ReadFile(filename)
	if err != nil {
		utils.Fatal(utils.NewError("CLI", 0, 0, fmt.Sprintf("system cannot find the file '%s'", filename)))
		os.Exit(1)
	}

	utils.Info(fmt.Sprintf("Running: %s", filename))
	fmt.Println()
	fmt.Printf("%sSource Code:%s\n%s\n", utils.ColorCyan, utils.ColorReset, string(source))

	// Step 1: Lexer 
	fmt.Printf("\n%s[LEXER OUTPUT]%s\n", utils.ColorGreen, utils.ColorReset)
	l := lexer.New(string(source))
	tokens := l.Tokenize()

	// Print lexical table
	utils.PrintLexicalTable(tokens)


	// Step 2: Parser
	p := parser.New(tokens)
    program := p.Parse()

	//log the AST
	fmt.Printf("\n%s[AST OUTPUT]%s\n", utils.ColorGreen, utils.ColorReset)
	fmt.Printf("%#v\n", program)
}
