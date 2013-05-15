/*

Main entry point to globol

*/

package main

import (
    "os"
    "fmt"
    "io"
    "bytes"
    "globol/lexer"
)

func main() {
    fmt.Println("buttslol %d", lexer.TOK_INDENT)

    var (
        file *os.File
        i int
        char byte
        atom_buffer bytes.Buffer
        ctx [33]int // 32 functions plus a string
        ctx_depth int
        /* err error */
        /* part []byte */
        /* prefix bool */
    )

    new_atom_buffer := func() {
        atom_buffer = *new(bytes.Buffer)
    }

    ctx_depth = 0
    ctx[ctx_depth] = lexer.CTX_MAIN

    buf := make([]byte, 1024)

    current_token := new(lexer.Token)

    look_back := func(n int) byte {
        /* TODO catch invalid indices */
        return atom_buffer.Bytes()[-n]
    }

    enter_ctx := func(c int) {
        ctx_depth++
        ctx[ctx_depth] = c
    }

    exit_ctx := func() int {
        var old int = ctx[ctx_depth]
        ctx_depth--
        return old
    }

    advance_token := func() {
        new_tok := new(lexer.Token)
        current_token.Next = new_tok
        current_token = new_tok
    }

    file, _ = os.Open(os.Args[1])
    for {
        n, err := file.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }

        // Iterate over bytes read
        i = 0
        for i < n {
            char = buf[i]
            if (ctx[ctx_depth] == lexer.CTX_STRING) {
                // Work out if we need out of this string
                if char == lexer.DELIM_STRING_END {
                    if look_back(1) == lexer.DELIM_STRING_END {
                        exit_ctx()
                        current_token.Type = lexer.TOK_STRING
                        current_token.Content = atom_buffer.Bytes()
                        advance_token()
                        new_atom_buffer()
                    }
                }
            } else if (false) {
                enter_ctx(lexer.CTX_STRING)
            }

            if (ctx[ctx_depth] != lexer.CTX_STRING) && lexer.IsAtomSeperator(char) {
                fmt.Println("hit an atom seperator")
                fmt.Println("Current state of atom_buffer:", atom_buffer)
                atom_buffer.Reset()
            } else {
                fmt.Println("> ", char)
                _ = atom_buffer.WriteByte(char)
            }
            i++
        }
        // Check if we've hit an atom seperator

    }

    /* contents,_ := ioutil.ReadFile(os.Args[1]) */

    /* TODO deal with errors */

    /* fmt.Println(contents) */


}
