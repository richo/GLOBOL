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

    current_token := token_list

    /* first_token := token_list */

    for {
        fmt.Println(string(current_token.Content))

        current_token = current_token.Next
        if (current_token == nil) {
            break
        }
    }
}
