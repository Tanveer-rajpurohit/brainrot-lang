# 🧠 BrainRot Lang

> *The no-cap programming language — where internet memes meet compiler design.*

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Mac%20%7C%20Linux-blue?style=flat)]()
[![Vibe](https://img.shields.io/badge/vibe-fr%20fr%20no%20cap-purple?style=flat)]()

BrainRot Lang is a fully working **interpreted programming language** built from scratch in Go. Every keyword is a real internet meme. It implements a complete compiler pipeline — Lexer → Parser → AST → Interpreter — covering all major phases of compiler design.

Built as a **Compiler Design mini project** (CE5023, Semester VI, Uka Tarsadia University).

---

## 🔥 What Does It Look Like?

```brainrot
let_him_cook fib(n) {
    chat_is_this_real (n <= 1) {
        take_this n
    }
    take_this fib(n - 1) + fib(n - 2)
}

let_him_cook main() {
    say_my_name("=== Fibonacci no cap ===")
    run_it_back (trust_me_bro i = 0; i < 8; i += 1) {
        say_my_name(fib(i))
    }
}
```

**Output:**
```
=== Fibonacci no cap ===
0
1
1
2
3
5
8
13
```

---

## 📖 Keyword Reference

| BrainRot Keyword     | Meaning           | Meme Origin              |
|----------------------|-------------------|--------------------------|
| `trust_me_bro`       | variable declare  | "trust me bro" Gen Z     |
| `say_my_name(...)`   | print             | Breaking Bad — Heisenberg|
| `chat_is_this_real`  | if                | "is this real?" caption  |
| `wait_hold_up`       | else if           | "wait, hold up" reaction |
| `nah_bro`            | else              | "nah bro" response       |
| `on_repeat`          | while loop        | on repeat cycle          |
| `run_it_back`        | for loop          | "run it back" sports     |
| `let_him_cook`       | function define   | "let him cook" 2023      |
| `take_this`          | return            | "take this" send-off     |
| `mission_abort`      | break             | abort mission            |
| `skip_this_one`      | continue          | "skip this one"          |
| `fr_fr`              | true              | Gen Z "for real"         |
| `cap`                | false             | cap/no cap meme          |
| `ghosted`            | null/nil          | dating meme              |

Operators `==`, `!=`, `+`, `-`, `*`, `/`, `%`, `**`, `++`, `--`, `&&`, `||`, `!` stay the same.

---

## 🚀 Installation

### Option 1 — Download Pre-built Binary (No Go needed)

1. Go to [Releases](../../releases) on GitHub
2. Download the binary for your OS:
   - `brainrot.exe` → Windows
   - `brainrot-linux` → Linux
   - `brainrot-mac` → Mac (Intel)
   - `brainrot-mac-arm` → Mac (Apple Silicon)
3. Add to PATH (see below)

**Windows — Add to PATH:**
```powershell
# Move brainrot.exe to a folder, e.g. C:\brainrot\
# Then add to PATH:
[System.Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";C:\brainrot", "User")
# Restart terminal, then:
brainrot run hello.brl
```

**Mac/Linux — Add to PATH:**
```bash
# Move binary to /usr/local/bin
sudo mv brainrot-linux /usr/local/bin/brainrot
sudo chmod +x /usr/local/bin/brainrot
# Now use from anywhere:
brainrot run hello.brl
```

---

### Option 2 — Build From Source (Go required)

```bash
# 1. Install Go from https://go.dev/dl/
# 2. Clone the repo
git clone https://github.com/Tanveer-rajpurohit/brainrot-lang.git
cd brainrot-lang

# 3. Build for your OS
go build -o brainrot main.go          # Mac/Linux
go build -o brainrot.exe main.go      # Windows

# 4. Run a program
./brainrot run examples/hello.brl
```

**Or run without building (development mode):**
```bash
go run main.go run examples/hello.brl
```

---

### Option 3 — Build All Platforms at Once

```powershell
# Windows (PowerShell)
powershell -ExecutionPolicy Bypass -File scripts/build.ps1
# Outputs: dist/brainrot.exe, dist/brainrot-linux, dist/brainrot-mac, dist/brainrot-mac-arm
```

---

## 📟 CLI Commands

```bash
brainrot run     <file.brl>   # Run a BrainRot program
brainrot tokens  <file.brl>   # Debug: show token stream (Lexer output)
brainrot ast     <file.brl>   # Debug: show AST tree (Parser output)
brainrot help                 # Show help
```

---

## 📁 Project Structure

```
brainrot-lang/
│
├── main.go                    ← CLI entry point
│
├── lexer/
│   ├── token.go               ← Token type definitions
│   └── lexer.go               ← DFA-based tokenizer (Unit II)
│
├── parser/
│   ├── ast.go                 ← AST node types
│   ├── parser.go              ← Recursive descent parser (Unit III)
│   ├── parseFunc.go           ← Statement parsing functions
│   └── parseExpr.go           ← Expression parsing chain
│
├── interpreter/
│   ├── interpreter.go         ← Interpreter struct + signals
│   ├── environment.go         ← Symbol table + scope chain (Unit V)
│   ├── evalStmt.go            ← Statement evaluator
│   ├── evalExpr.go            ← Expression evaluator
│   └── helpers.go             ← isTruthy, evalInfix, formatValue
│
├── utils/
│   ├── errors.go              ← Error types
│   ├── logger.go              ← Colored terminal output
│   ├── printer.go             ← AST printer + token table
│   └── utils.go               ← Token category helper
│
├── examples/
│   ├── hello.brl              ← Variables, print, booleans
│   ├── ifelse.brl             ← Conditionals + else-if chains
│   ├── for.brl                ← For loops
│   ├── fizzbuzz.brl           ← Classic FizzBuzz
│   └── fibonacci.brl          ← Recursion + functions
│
├── scripts/
│   ├── build.ps1              ← Cross-platform build (Windows)
│   └── test.ps1               ← Automated test runner
│
├── dist/                      ← Pre-built binaries (gitignored)
├── go.mod
├── LICENSE
└── README.md
```

---

## 📜 Example Programs

### Hello World
```brainrot
let_him_cook main() {
    trust_me_bro name = "Walter White"
    say_my_name("my name is " + name)
}
```

### FizzBuzz
```brainrot
let_him_cook main() {
    run_it_back (trust_me_bro i = 1; i <= 20; i += 1) {
        chat_is_this_real (i % 15 == 0) {
            say_my_name("FizzBuzz no cap")
        }
        wait_hold_up (i % 3 == 0) {
            say_my_name("Fizz fr fr")
        }
        wait_hold_up (i % 5 == 0) {
            say_my_name("Buzz W")
        }
        nah_bro {
            say_my_name(i)
        }
    }
}
```

### Functions + Recursion
```brainrot
let_him_cook factorial(n) {
    chat_is_this_real (n <= 1) {
        take_this 1
    }
    take_this n * factorial(n - 1)
}

let_him_cook main() {
    say_my_name(factorial(5))
}
```

---

## 🏗️ Compiler Pipeline

```
Source Code (.brl file)
        │
        ▼
┌─────────────┐
│    LEXER    │  Tokenizes raw text → stream of tokens
│  (Unit II)  │  DFA-based scanner, recognizes all keywords
└──────┬──────┘
       │  []Token
       ▼
┌─────────────┐
│   PARSER    │  Builds Abstract Syntax Tree from tokens
│  (Unit III) │  Recursive descent, LL(1) parsing
└──────┬──────┘
       │  AST
       ▼
┌─────────────────┐
│  INTERPRETER    │  Walks AST and executes the program
│   (Unit V)      │  Symbol tables, scope chains, closures
└─────────────────┘
       │
       ▼
   Program Output
```

---

## 🎓 Syllabus Coverage

| Syllabus Unit | Topic | Implemented In |
|---|---|---|
| Unit I | Compiler phases, types, tools | Entire pipeline |
| Unit II | Lexical analysis, DFA, tokens | `lexer/lexer.go` |
| Unit III | CFG, recursive descent, LL(1), AST | `parser/` |
| Unit V | Symbol tables, scope, runtime env | `interpreter/environment.go` |

---

## 🛠️ Built With

Go 1.22+ — zero external dependencies, pure standard library.

---

## 📄 License

MIT — do whatever you want, no cap.

## ⭐ Give it a star if you fw it!