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
)

const (
    CTX_NEWLINE = iota // Always first, treated as a mask
    CTX_MAIN = iota
    CTX_FUNC = iota
    CTX_STRING = iota
)

func IsAtomSeperator(c byte) bool {
    var (
        ret bool
    )
    ret = c == 32 /* space */ ||
          c == 44 /* comma */ ||
          c == 10 /* newline */

    return ret
}
