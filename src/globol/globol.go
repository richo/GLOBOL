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

const CONTEXT_STACK_DEPTH = 33

func main() {

    var (
        file *os.File
        i int
        char byte
        atom_buffer bytes.Buffer
        ctx [CONTEXT_STACK_DEPTH]int // 32 functions plus a string
        ctx_depth int
        atom_idx int = 0
        /* err error */
        /* part []byte */
        /* prefix bool */
    )

    new_atom_buffer := func() {
        atom_buffer = *new(bytes.Buffer)
        atom_idx = 0
    }

    ctx_depth = 0
    ctx[ctx_depth] = lexer.CTX_MAIN

    buf := make([]byte, 1024)

    current_token := new(lexer.Token)

    look_back := func(n int) byte {
        if (atom_idx == 0) {
            return 0 // XXX
        }
        return atom_buffer.Bytes()[atom_idx - n]
    }

    add_to_buf := func(n byte) {
        _ = atom_buffer.WriteByte(n)
        atom_idx++
    }

    enter_ctx := func(c int) {
        /* assert(ctx_depth < CONTEXT_STACK_DEPTH) */
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

    push_atom := func(tok_type int, content []byte, reset bool) {
        fmt.Println("Parsed Token:", tok_type, string(content))
        current_token.Type = tok_type
        current_token.Content = content
        advance_token()
        if (reset) {
            new_atom_buffer()
        }
    }

    start_of_line := func() bool {
        return ctx[ctx_depth] & lexer.CTX_NEWLINE == lexer.CTX_NEWLINE
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
            i++
            if (ctx[ctx_depth] == lexer.CTX_STRING) {
                // Work out if we need out of this string
                if char == lexer.DELIM_STRING_END &&
                    look_back(1) == lexer.DELIM_STRING_END {
                        // Drop the last character from our buffer
                        _ = atom_buffer.Next(1)
                        exit_ctx()
                        push_atom(lexer.TOK_STRING, atom_buffer.Bytes()[:atom_buffer.Len()-1], true)
                } else { // Just an ordinary char in a string
                    add_to_buf(char)
                }
                continue
            }

            if char == lexer.DELIM_STRING_BEGIN {
                if look_back(1) == lexer.DELIM_STRING_BEGIN {
                    enter_ctx(lexer.CTX_STRING)
                    continue
                }
            }

            /** END STRING HANDLING **/

            if (lexer.IsAtomSeperator(char)) {
                if !(char == lexer.MARK_INDENT && start_of_line()) {
                    if (atom_buffer.Len() > 0) {
                        push_atom(lexer.TOK_ATOM, atom_buffer.Bytes(), true)
                    }
                    if (char == lexer.MARK_NEWLINE) {
                        ctx[ctx_depth] = ctx[ctx_depth] | lexer.CTX_NEWLINE
                        push_atom(lexer.TOK_ENDSTATEMENT, []byte(";;"), false)
                    }
                    continue
                }
            }

            /* Safe to assume we're not in a string */

            if (start_of_line()) {
                if (char == lexer.MARK_INDENT) {
                    if (look_back(1) == lexer.MARK_INDENT) {
                        push_atom(lexer.TOK_INDENT, []byte("||"), true)
                        continue
                    }
                } else {
                    ctx[ctx_depth] = ctx[ctx_depth] & ^lexer.CTX_NEWLINE
                    // Not an indent, start of line, drop out of NL context
                }
            }

            add_to_buf(char)

        }
        // Check if we've hit an atom seperator

    }

    /* contents,_ := ioutil.ReadFile(os.Args[1]) */

    /* TODO deal with errors */

    /* fmt.Println(contents) */


}
