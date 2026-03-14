# Brainrot Lang

A meme‑inspired toy programming language built in **Go**.

Brainrot Lang replaces traditional programming keywords with internet
brainrot slang like `say_my_name`, `chat_is_this_real`, and
`run_it_back`.

This project demonstrates how a programming language works internally by
implementing:

-   Lexer
-   Parser
-   AST (Abstract Syntax Tree)
-   Interpreter

------------------------------------------------------------------------

#  Features

-   Meme‑based syntax
-   Simple interpreter written in Go
-   Supports:
    -   Variables
    -   Conditionals
    -   Loops
    -   Functions
    -   Return values
-   CLI runner for `.brl` files

------------------------------------------------------------------------

#  Brainrot Keywords

  Brainrot Keyword    Meaning
  ------------------- ----------
  say_my_name         print
  let_him_cook        function
  trust_me_bro        variable
  chat_is_this_real   if
  wait_hold_up        else if
  nah_bro             else
  on_repeat           while
  run_it_back         for
  take_this           return
  mission_abort       break
  skip_this_one       continue
  fr_fr               true
  cap                 false
  ghosted             nil

------------------------------------------------------------------------

# 🚀 Quick Start

### Run during development

``` bash
go run main.go run examples/hello.brl
```

------------------------------------------------------------------------

### Build CLI

``` bash
go build -o brainrot main.go
```

Run program:

``` bash
./brainrot run examples/hello.brl
```

------------------------------------------------------------------------

# 🖥 Cross Platform Build

### Windows

``` bash
GOOS=windows GOARCH=amd64 go build -o brainrot.exe main.go
```

### Mac

``` bash
GOOS=darwin GOARCH=amd64 go build -o brainrot main.go
```

### Linux

``` bash
GOOS=linux GOARCH=amd64 go build -o brainrot main.go
```

------------------------------------------------------------------------

# 📂 Project Structure

    brainrot-lang
    │
    ├── examples
    │   ├── hello.brl
    │   ├── fizzbuzz.brl
    │   └── fibonacci.brl
    │
    ├── lexer
    │   ├── lexer.go
    │   └── token.go
    │
    ├── parser
    │   ├── parser.go
    │   └── ast.go
    │
    ├── interpreter
    │
    ├── utils
    │   ├── errors.go
    │   ├── logger.go
    │   └── utils.go
    │
    ├── main.go
    ├── go.mod
    └── README.md

------------------------------------------------------------------------

# 📜 Example Program

``` brainrot
let_him_cook main(){

    trust_me_bro x = 5

    chat_is_this_real (x > 3){
        say_my_name("hello world fr fr")
    } nah_bro {
        say_my_name("cap")
    }

}
```

------------------------------------------------------------------------

# 🔁 Loop Example

``` brainrot
run_it_back (i = 0; i < 5; i = i + 1){
    say_my_name(i)
}
```

------------------------------------------------------------------------

# 🎯 Goals of this Project

This project was built to:

-   Learn compiler / interpreter architecture
-   Practice Go language
-   Understand lexing and parsing
-   Build a fun programming language

------------------------------------------------------------------------

# 🛠 Future Improvements

-   Type system
-   Standard library
-   Better error messages
-   REPL mode
-   Bytecode VM

------------------------------------------------------------------------

# 📄 License

MIT License

------------------------------------------------------------------------

# ⭐ If you like this project

Give it a star on GitHub!
