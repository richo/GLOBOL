package lexer

type Token struct {
    Type int
    Content []byte
    Next *Token
}

/* type LexStackV */
