# 🧠 BrainRot Lang

> *The no-cap programming language — where internet memes meet compiler design.*

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Mac%20%7C%20Linux-blue?style=flat)]()
[![Vibe](https://img.shields.io/badge/vibe-fr%20fr%20no%20cap-purple?style=flat)]()

BrainRot Lang is a fully working **interpreted programming language** built from scratch in Go. Every keyword is a real internet meme. It implements a complete compiler pipeline — Lexer → Parser → AST → Interpreter — covering all major phases of compiler design.

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

| BrainRot Keyword     | Meaning           | Meme Origin               |
|----------------------|-------------------|---------------------------|
| `trust_me_bro`       | variable declare  | "trust me bro" Gen Z      |
| `say_my_name(...)`   | print             | Breaking Bad — Heisenberg |
| `chat_is_this_real`  | if                | "is this real?" caption   |
| `wait_hold_up`       | else if           | "wait, hold up" reaction  |
| `nah_bro`            | else              | "nah bro" response        |
| `on_repeat`          | while loop        | on repeat cycle           |
| `run_it_back`        | for loop          | "run it back" sports      |
| `let_him_cook`       | function define   | "let him cook" 2023       |
| `take_this`          | return            | "take this" send-off      |
| `mission_abort`      | break             | abort mission             |
| `skip_this_one`      | continue          | "skip this one"           |
| `fr_fr`              | true              | Gen Z "for real for real" |
| `cap`                | false             | cap/no cap meme           |
| `ghosted`            | null/nil          | dating meme               |

Operators `==`, `!=`, `+`, `-`, `*`, `/`, `%`, `**`, `++`, `--`, `&&`, `||`, `!` stay the same as normal.

---

## 🚀 Installation

> **No Go installation needed.** Just download the binary for your OS and add it to PATH. That's it.

---

### 🪟 Windows

#### Step 1 — Download the binary
Go to [Releases](../../releases) and download **`brainrot.exe`**

#### Step 2 — Create install folder and move binary

Open **PowerShell** and run:
```powershell
mkdir C:\Users\$env:USERNAME\brainrot
move brainrot.exe C:\Users\$env:USERNAME\brainrot\
```

This creates `C:\Users\YourName\brainrot\` and moves the binary there.

#### Step 3 — Add to PATH

**PowerShell method (recommended):**
```powershell
[System.Environment]::SetEnvironmentVariable(
    "PATH",
    $env:PATH + ";C:\Users\$env:USERNAME\brainrot",
    "User"
)
```

**Manual (GUI method):**
1. Press `Windows + S` → search **"Environment Variables"**
2. Click **Edit the system environment variables**
3. Click the **Environment Variables** button
4. Under **User variables** → click `Path` → click **Edit**
5. Click **New** → type `C:\Users\YourName\brainrot` (replace `YourName` with your actual Windows username)
6. Click **OK** → **OK** → **OK**

#### Step 4 — Restart terminal and verify
**Close PowerShell/terminal completely and reopen it**, then:
```powershell
brainrot help
```

You should see the help menu. ✅

#### Step 5 — Run your first program
```powershell
brainrot run examples/hello.brt
brainrot run examples/fibonacci.brt
brainrot run examples/fizzbuzz.brt
```

---

### 🍎 Mac (Intel + Apple Silicon)

#### Step 1 — Download the binary
Go to [Releases](../../releases) and download:
- **`brainrot-mac`** → Mac Intel (x86_64)
- **`brainrot-mac-arm`** → Mac Apple Silicon (M1 / M2 / M3)

#### Step 2 — Open Terminal and install

**Intel Mac:**
```bash
sudo mv ~/Downloads/brainrot-mac /usr/local/bin/brainrot
sudo chmod +x /usr/local/bin/brainrot
```

**Apple Silicon (M1 / M2 / M3):**
```bash
sudo mv ~/Downloads/brainrot-mac-arm /usr/local/bin/brainrot
sudo chmod +x /usr/local/bin/brainrot
```

#### Step 3 — Allow Mac to run it (first time only)

Mac blocks downloaded binaries by default. Run this once:
```bash
xattr -d com.apple.quarantine /usr/local/bin/brainrot
```

If you see a popup instead — go to **System Settings → Privacy & Security → General** and click **Allow Anyway**.

#### Step 4 — Verify
```bash
brainrot help
```

You should see the help menu. ✅

#### Step 5 — Run your first program
```bash
brainrot run examples/hello.brt
brainrot run examples/fibonacci.brt
brainrot run examples/fizzbuzz.brt
```

---

### 🐧 Linux

#### Step 1 — Download the binary
Go to [Releases](../../releases) and download **`brainrot-linux`**

#### Step 2 — Open terminal and install
```bash
sudo mv ~/Downloads/brainrot-linux /usr/local/bin/brainrot
sudo chmod +x /usr/local/bin/brainrot
```

#### Step 3 — Verify
```bash
brainrot help
```

You should see the help menu. ✅

#### Step 4 — Run your first program
```bash
brainrot run examples/hello.brt
brainrot run examples/fibonacci.brt
brainrot run examples/fizzbuzz.brt
```

---

## 📟 CLI Commands

```bash
brainrot run     <file.brt>   # Run a BrainRot program
brainrot tokens  <file.brt>   # Debug: show token stream (Lexer output)
brainrot ast     <file.brt>   # Debug: show AST tree (Parser output)
brainrot help                 # Show help
```

**Examples:**
```bash
brainrot run examples/hello.brt
brainrot run examples/ifelse.brt
brainrot run examples/for.brt
brainrot run examples/fizzbuzz.brt
brainrot run examples/fibonacci.brt

# Debug modes — great for understanding how the compiler works
brainrot tokens examples/hello.brt
brainrot ast examples/hello.brt
```

---

## 🛠️ Build From Source (For Developers)

> Only needed if you want to modify the language itself. Regular users skip this.

**Requires:** Go 1.22+ from [https://go.dev/dl/](https://go.dev/dl/)

```bash
# Clone
git clone https://github.com/Tanveer-rajpurohit/brainrot-lang.git
cd brainrot-lang

# Build for your current OS
go build -o brainrot.exe main.go     # Windows
go build -o brainrot main.go         # Mac / Linux

# Build for ALL platforms at once (Windows PowerShell)
powershell -ExecutionPolicy Bypass -File scripts/build.ps1

# Run without building (dev mode)
go run main.go run examples/hello.brt
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
│   └── lexer.go               ← DFA-based tokenizer
│
├── parser/
│   ├── ast.go                 ← AST node types
│   ├── parser.go              ← Recursive descent parser
│   ├── parseFunc.go           ← Statement parsing functions
│   └── parseExpr.go           ← Expression parsing + precedence chain
│
├── interpreter/
│   ├── interpreter.go         ← Interpreter struct + signals
│   ├── environment.go         ← Symbol table + scope chain
│   ├── evalStmt.go            ← Statement evaluator + program entry
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
│   ├── hello.brt              ← Variables, print, booleans
│   ├── ifelse.brt             ← Conditionals + else-if chains
│   ├── for.brt                ← For loops
│   ├── fizzbuzz.brt           ← Classic FizzBuzz
│   └── fibonacci.brt          ← Recursion + functions
│
├── scripts/
│   ├── build.ps1              ← Cross-platform build script
│   └── test.ps1               ← Automated test runner
│
├── dist/                      ← Pre-built binaries
├── Makefile                   ← Dev shortcuts (make run, make test)
├── go.mod
├── LICENSE
├── PROJECT.md                 ← Full technical project documentation
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

### If / Else If / Else
```brainrot
let_him_cook main() {
    trust_me_bro score = 85
    chat_is_this_real (score >= 90) {
        say_my_name("S tier no cap")
    }
    wait_hold_up (score >= 75) {
        say_my_name("pretty mid ngl")
    }
    nah_bro {
        say_my_name("ur cooked bro")
    }
}
```

### For Loop
```brainrot
let_him_cook main() {
    run_it_back (trust_me_bro i = 0; i < 5; i += 1) {
        say_my_name("grind " + i)
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
Source Code (.brt file)
        │
        ▼
┌─────────────┐
│    LEXER    │  Tokenizes raw text → stream of tokens
│             │  DFA-based scanner, recognizes all keywords
└──────┬──────┘
       │  []Token
       ▼
┌─────────────┐
│   PARSER    │  Builds Abstract Syntax Tree from tokens
│             │  Recursive descent, LL(1) parsing
└──────┬──────┘
       │  AST
       ▼
┌─────────────────┐
│  INTERPRETER    │  Walks AST and executes the program
│                 │  Symbol tables, scope chains, closures
└─────────────────┘
       │
       ▼
   Program Output
```

---

## 📄 License

MIT — do whatever you want, no cap.

---

## ⭐ Give it a star if you fw it!