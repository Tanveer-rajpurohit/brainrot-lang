# BrainRot Lang — Project Documentation

## 1. What Is BrainRot Lang?

BrainRot Lang is a fully functional interpreted programming language built from scratch in Go. It replaces traditional programming keywords with internet meme slang — `trust_me_bro` instead of `var`, `chat_is_this_real` instead of `if`, `let_him_cook` instead of `function`. Despite the meme syntax, it is a complete, working language that demonstrates every major phase of compiler design covered in CE5023.

**It is an interpreter language.** This means source code is not compiled to machine code. Instead, after the source is lexed and parsed into an AST, a tree-walking interpreter executes the AST directly in memory. This is the same approach used by early Python and Ruby.

---

## 2. Why Interpreter and Not Compiler?

| | Interpreter (BrainRot) | Compiler (like GCC) |
|---|---|---|
| Front end (Lexer + Parser + AST) | ✅ Identical | ✅ Identical |
| Back end | Tree-walker executes AST directly | Code generator → machine code |
| Output | Terminal output directly | `.exe` / binary file |
| Difficulty | Achievable in 5 days | Requires months |
| Used by | Python, Ruby, early JS | C, Go, Rust |

---

## 3. Compiler Pipeline — Phase by Phase

### Phase 1: Lexical Analysis (Unit II)

**File:** `lexer/lexer.go`, `lexer/token.go`

The Lexer reads raw source code character by character and converts it into a flat list of tokens. This is a practical implementation of the DFA (Deterministic Finite Automaton) theory from Unit II.

```
Input:  trust_me_bro x = 42
Output: [VAR:"trust_me_bro"] [IDENT:"x"] [ASSIGN:"="] [INT:"42"] [EOF]
```

Key concepts implemented:
- DFA-based character scanning (`nextToken()`)
- Input buffering with two-pointer technique (`position`, `readPosition`)
- Keyword recognition via lookup table (`LookupIdent()`)
- Multi-character token handling with peek (`peekChar()`) — e.g. `==` vs `=`, `+=` vs `+`
- Comment skipping (`#` single-line comments)
- Escape sequence handling in strings (`\n`, `\t`, `\\`)

### Phase 2: Syntax Analysis / Parsing (Unit III)

**File:** `parser/parser.go`, `parser/parseFunc.go`, `parser/parseExpr.go`

The Parser reads the token stream and builds an Abstract Syntax Tree (AST) — a tree structure that captures the meaning and structure of the program.

```
Tokens: [VAR] [IDENT:"x"] [ASSIGN] [INT:"42"]
AST:    VarStatement { Name: "x", Value: IntegerLiteral{42} }
```

Key concepts implemented:
- Context-Free Grammar (CFG) — each grammar rule becomes one function
- Recursive Descent Parsing — `parseStatement()` calls `parseIfStatement()` which calls `parseExpression()` which calls `parseTerm()` etc.
- LL(1) parsing — one token lookahead using `peek()`
- Operator precedence via function call chain (lower in chain = higher precedence)
- Error recovery — collects all errors instead of stopping at first one
- Abstract Syntax Tree with 20+ node types

### Phase 3: AST Construction

**File:** `parser/ast.go`

The AST uses Go interfaces for type safety:
- `Expression` interface — nodes that produce a value (numbers, operators, function calls)
- `Statement` interface — nodes that perform actions (print, if, loop, assign)
- Embedding pattern reduces boilerplate (each node embeds `stmtNode` or `exprNode`)

### Phase 4: Semantic Analysis + Runtime (Unit V)

**File:** `interpreter/environment.go`, `interpreter/evalStmt.go`, `interpreter/evalExpr.go`

The Interpreter walks the AST and executes each node. The Environment (symbol table) stores variables in a scope chain.

Key concepts implemented:
- Symbol Table — `Environment.store` (map of variable names to values)
- Scope Chain — each function call creates a child environment with a pointer to the parent
- Closure — functions capture their definition-time environment
- Control flow signals — `ReturnValue`, `BreakSignal`, `ContinueSignal` bubble up the call stack

---

## 4. Language Specification

### Program Structure

Every BrainRot program must have exactly one `let_him_cook main()` function as the entry point. Global variables and function definitions are allowed outside `main()`. All executable code must be inside a function.

```
program → globalDecl* main_func

globalDecl → varDecl | funcDecl

varDecl  → "trust_me_bro" IDENT "=" expression
funcDecl → "let_him_cook" IDENT "(" params? ")" "{" statement* "}"
```

### Grammar (Context-Free Grammar)

```
statement      → varDecl | assignStmt | printStmt | ifStmt
               | whileStmt | forStmt | funcDecl | returnStmt
               | breakStmt | continueStmt | exprStmt

varDecl        → "trust_me_bro" IDENT "=" expression
assignStmt     → IDENT ("=" | "+=" | "-=" | "*=" | "/=") expression
printStmt      → "say_my_name" "(" expression ")"
ifStmt         → "chat_is_this_real" expression "{" statement* "}"
                 ("wait_hold_up" expression "{" statement* "}")*
                 ("nah_bro" "{" statement* "}")?
whileStmt      → "on_repeat" expression "{" statement* "}"
forStmt        → "run_it_back" "(" varDecl ";" expression ";" assignStmt ")"
                 "{" statement* "}"
funcDecl       → "let_him_cook" IDENT "(" params? ")" "{" statement* "}"
returnStmt     → "take_this" expression?
breakStmt      → "mission_abort"
continueStmt   → "skip_this_one"

expression     → or_expr
or_expr        → and_expr ("||" and_expr)*
and_expr       → equality ("&&" equality)*
equality       → comparison (("==" | "!=") comparison)*
comparison     → term (("<" | ">" | "<=" | ">=") term)*
term           → factor (("+" | "-") factor)*
factor         → unary (("*" | "/" | "%" | "**") unary)*
unary          → ("!" | "-") unary | postfix
postfix        → primary ("++" | "--")?
primary        → INT | FLOAT | STRING | "fr_fr" | "cap" | "ghosted"
               | IDENT | IDENT "(" arguments? ")" | "(" expression ")"
               | "[" elements? "]"
```

### Data Types

| Type | BrainRot | Example |
|---|---|---|
| Integer | `INT` | `42`, `-10` |
| Float | `FLOAT` | `3.14` |
| String | `STRING` | `"hello"` |
| Boolean | `fr_fr` / `cap` | `fr_fr` = true |
| Nil | `ghosted` | `ghosted` |
| Array | `[...]` | `[1, 2, 3]` |
| Function | `let_him_cook` | first-class value |

---

## 5. Error Handling

BrainRot Lang uses **error recovery** — the parser collects all errors instead of stopping at the first one. This is the approach used by professional compilers (GCC, Go compiler) rather than Python's stop-on-first-error approach.

Three categories of errors:
- **Lexer errors** — illegal characters
- **Parser errors** — syntax violations, missing tokens
- **Runtime errors** — undefined variables, type mismatches, division by zero

---

## 6. How to Use in Demo

```bash
# Run a program
go run main.go run examples/fibonacci.brt

# Show token stream (demonstrates Lexer)
go run main.go tokens examples/hello.brt

# Show AST (demonstrates Parser)
go run main.go ast examples/hello.brt

# Show all three in sequence for demo
go run main.go tokens examples/hello.brt
go run main.go ast examples/hello.brt
go run main.go run examples/hello.brt
```

---

## 7. Files Quick Reference

| File | Lines | Purpose |
|---|---|---|
| `lexer/lexer.go` | ~250 | DFA tokenizer — Unit II core |
| `lexer/token.go` | ~100 | All token types and keyword map |
| `parser/ast.go` | ~150 | All 20+ AST node types |
| `parser/parser.go` | ~100 | Parser core + helper methods |
| `parser/parseExpr.go` | ~200 | Expression chain (precedence climbing) |
| `parser/parseFunc.go` | ~200 | Statement parsing functions |
| `interpreter/environment.go` | ~60 | Symbol table — Unit V core |
| `interpreter/evalStmt.go` | ~200 | Statement execution + program entry |
| `interpreter/evalExpr.go` | ~100 | Expression evaluation |
| `interpreter/helpers.go` | ~150 | Math, comparison, formatting |
| `interpreter/interpreter.go` | ~40 | Interpreter struct + error handling |

---

## 8. Future Extensions

- **Type system** — static type checking before execution
- **Standard library** — math, string, array functions built-in
- **REPL** — interactive `brainrot repl` mode
- **Bytecode compilation** — compile AST to bytecode + BrainRot VM
- **Go transpiler** — `brainrot compile` flag generates `.go` source
- **VS Code extension** — full syntax highlighting + Ctrl+/ comments

---

## 9. References

- A.V. Aho, Monica Lam, Ravi Sethi, J.D. Ullman — *Compilers: Principles, Techniques & Tools* (Dragon Book)
- Robert Nystrom — *Crafting Interpreters* (https://craftinginterpreters.com) — free online
- GeeksForGeeks — Lexical Analysis, Syntax Analysis, Symbol Tables