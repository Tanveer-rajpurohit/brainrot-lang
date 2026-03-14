package lexer

// Token types
type TokenType string

const (
	// Special tokens
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
	NEWLINE TokenType = "NEWLINE"

	// Literals
	IDENT  TokenType = "IDENT"  // identifier
	INT    TokenType = "INT"    // integer
	FLOAT  TokenType = "FLOAT"  // float
	STRING TokenType = "STRING" // string
	// ── BrainRot Keywords ──────────────────────
	//  trust_me_bro   → var       ("trust me bro")
	//  say_my_name    → print     (Breaking Bad - Heisenberg)
	//  chat_is_this_real → if     (is this real? caption)
	//  nah_bro        → else      ("nah bro" response)
	//  on_repeat      → while     (on repeat cycle)
	//  run_it_back    → for       (run it back loop)
	//  let_him_cook   → func      ("let him cook" 2023)
	//  take_this      → return    ("take this" send off)
	//  fr_fr          → true      (Gen Z no cap)
	//  cap            → false     (cap/no cap)
	//  ghosted        → nil/null  (dating meme)
	//  mission_abort  → break     (abort mission)
	//  skip_this_one  → continue  ("skip this one")

	VAR      TokenType = "trust_me_bro"
	PRINT    TokenType = "say_my_name"
	IF       TokenType = "chat_is_this_real"
	ELSE     TokenType = "nah_bro"
	WHILE    TokenType = "on_repeat"
	FOR      TokenType = "run_it_back"
	FUNC     TokenType = "let_him_cook"
	RETURN   TokenType = "take_this"
	TRUE     TokenType = "fr_fr"
	FALSE    TokenType = "cap"
	NIL      TokenType = "ghosted"
	BREAK    TokenType = "mission_abort"
	CONTINUE TokenType = "skip_this_one"

	// ── Arithmetic Operators ───────────────────
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	PERCENT  TokenType = "%"
	POWER    TokenType = "**"

	// ── Comparison Operators ───────────────────
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="
	LT     TokenType = "<"
	GT     TokenType = ">"
	LTE    TokenType = "<="
	GTE    TokenType = ">="

	// ── Logical Symbol Operators ───────────────
	AND TokenType = "&&"
	OR  TokenType = "||"
	NOT TokenType = "!"

	// ── Assignment ─────────────────────────────
	ASSIGN       TokenType = "="
	PLUS_ASSIGN  TokenType = "+="
	MINUS_ASSIGN TokenType = "-="

	// ── Delimiters ─────────────────────────────
	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"

	// ── Punctuation ────────────────────────────
	SEMICOLON TokenType = ";"
	COMMA     TokenType = ","
	DOT       TokenType = "."
	COLON     TokenType = ":"
	ARROW     TokenType = "=>"
)

// Token represents a single token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Keywords map
var keywords = map[string]TokenType{
	"trust_me_bro":      VAR,
	"say_my_name":       PRINT,
	"chat_is_this_real": IF,
	"nah_bro":           ELSE,
	"on_repeat":         WHILE,
	"run_it_back":       FOR,
	"let_him_cook":      FUNC,
	"take_this":         RETURN,
	"fr_fr":             TRUE,
	"cap":               FALSE,
	"ghosted":           NIL,
	"mission_abort":     BREAK,
	"skip_this_one":     CONTINUE,
}

// LookupIdent checks if a word is a BrainRot keyword or a variable name
func LookupIdent(ident string) TokenType {
	if tok, exists := keywords[ident]; exists {
		return tok
	}
	return IDENT
}
