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
	//  ngl          → var       ("not gonna lie")
	//  say_my_name  → print     (Breaking Bad - Heisenberg)
	//  bet          → if        (TikTok/Gen Z)
	//  nope         → else      (internet classic)
	//  rizz_while   → while     (Rizzler meme 2024)
	//  grind_for    → for       (hustle culture)
	//  let_him_cook → func      ("let him cook" 2023)
	//  served       → return    ("she served")
	//  fr_fr        → true      (Gen Z no cap)
	//  cap          → false     (cap/no cap)
	//  ghosted      → nil/null  (dating meme)
	//  slay         → break     (slay meme)
	//  pushing      → continue  ("keep pushing")
	//  no_cap       → &&        (no cap slang)
	//  or_what      → ||        ("or what??" meme)

	VAR         TokenType = "ngl"
	PRINT       TokenType = "say_my_name"
	IF          TokenType = "bet"
	ELSE        TokenType = "nope"
	WHILE       TokenType = "rizz_while"
	FOR         TokenType = "grind_for"
	FUNC        TokenType = "let_him_cook"
	RETURN      TokenType = "served"
	TRUE        TokenType = "fr_fr"
	FALSE       TokenType = "cap"
	NIL         TokenType = "ghosted"
	BREAK       TokenType = "slay"
	CONTINUE    TokenType = "pushing"
	LOGICAL_AND TokenType = "no_cap"
	LOGICAL_OR  TokenType = "or_what"

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
	"ngl":         VAR,
	"say_my_name": PRINT,
	"bet":         IF,
	"nope":        ELSE,

	"rizz_while":   WHILE,
	"grind_for":    FOR,
	"let_him_cook": FUNC,
	"served":       RETURN,
	"fr_fr":        TRUE,
	"cap":          FALSE,
	"ghosted":      NIL,
	"slay":         BREAK,
	"pushing":      CONTINUE,
	"no_cap":       LOGICAL_AND,
	"or_what":      LOGICAL_OR,
}

// LookupIdent checks if a word is a BrainRot keyword or a variable name
func LookupIdent(ident string) TokenType {
	if tok, exists := keywords[ident]; exists {
		return tok
	}
	return IDENT
}
