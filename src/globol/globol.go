/*

Main entry point to globol

*/

package main

import (
    "os"
    "fmt"
    "globol/lexer"
)

func main() {
    var (
        file *os.File
        token_list *lexer.Token
    )
    file, _ = os.Open(os.Args[1])

    token_list = lexer.Lex(file)
    fmt.Println(token_list)
}
