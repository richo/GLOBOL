package lexer

const (
 	 DELIM_STRING_BEGIN = 96
 	 DELIM_STRING_END = 39
 	 FOO = iota
)

const (
 	 TOK_ATOM = iota
 	 TOK_INDENT = iota
 	 TOK_COMMENT = iota
 	 TOK_STRING = iota
 	 TOK_ENDSTATEMENT = iota
)

const (
 	 _ = iota // Burn zero, or you'll wind up looking like an idiot
 	 CTX_NEWLINE = iota // Always first, treated as a mask
 	 CTX_MAIN = iota
 	 CTX_FUNC = iota
 	 CTX_STRING = iota
)

const (
 	 MARK_NEWLINE = 10
 	 MARK_SPACE = 32
 	 MARK_INDENT = MARK_SPACE
 	 MARK_COMMA = 44
)

func IsAtomSeperator(c byte) bool {
 	 var (
 	  	 ret bool
 	 )

 	 ret = c == MARK_SPACE ||
 	  	   c == MARK_COMMA ||
 	  	   c == MARK_NEWLINE

 	 return ret
}
