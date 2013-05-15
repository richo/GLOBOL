package lexer

const (
    TOK_INDENT = iota
    TOK_COMMENT = iota
)

func IsAtomSeperator(c byte) bool {
    var (
        ret bool
    )
    ret = c == 32 /* space */ ||
          c == 44 /* comma */

    return ret
}
