package lexer

import (
    "io"
    "os"
    "bytes"
)

type Token struct {
    Type int
    Content []byte
    Next *Token
}

const CONTEXT_STACK_DEPTH = 33

func Lex(file *os.File) *Token {

    var (
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
    ctx[ctx_depth] = CTX_MAIN

    buf := make([]byte, 1024)

    current_token := new(Token)
    first_token := current_token

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
        new_tok := new(Token)
        new_tok.Next = nil
        current_token.Next = new_tok
        current_token = new_tok
    }

    push_atom := func(tok_type int, content []byte, reset bool) {
        current_token.Type = tok_type
        /* dup(2) ... how does it work in go? */
        current_token.Content = []byte(string(content))
        advance_token()
        if (reset) {
            new_atom_buffer()
        }
    }

    start_of_line := func() bool {
        return ctx[ctx_depth] & CTX_NEWLINE == CTX_NEWLINE
    }

    for {
        n, err := file.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }

        // Iterate over bytes read
        i = 0
        for i < n {
            char = buf[i]
            i++
            if (ctx[ctx_depth] == CTX_STRING) {
                // Work out if we need out of this string
                if char == DELIM_STRING_END &&
                    look_back(1) == DELIM_STRING_END {
                        // Drop the last character from our buffer
                        _ = atom_buffer.Next(1)
                        exit_ctx()
                        push_atom(TOK_STRING, atom_buffer.Bytes()[:atom_buffer.Len()-1], true)
                } else { // Just an ordinary char in a string
                    add_to_buf(char)
                }
                continue
            }

            if char == DELIM_STRING_BEGIN {
                if look_back(1) == DELIM_STRING_BEGIN {
                    enter_ctx(CTX_STRING)
                    continue
                }
            }

            /** END STRING HANDLING **/

            if (IsAtomSeperator(char)) {
                if !(char == MARK_INDENT && start_of_line()) {
                    if (atom_buffer.Len() > 0) {
                        push_atom(TOK_ATOM, atom_buffer.Bytes(), true)
                    }
                    if (char == MARK_NEWLINE) {
                        ctx[ctx_depth] = ctx[ctx_depth] | CTX_NEWLINE
                        push_atom(TOK_ENDSTATEMENT, []byte(";;"), false)
                    }
                    continue
                }
            }

            /* Safe to assume we're not in a string */

            if (start_of_line()) {
                if (char == MARK_INDENT) {
                    if (look_back(1) == MARK_INDENT) {
                        push_atom(TOK_INDENT, []byte("||"), true)
                        continue
                    }
                } else {
                    ctx[ctx_depth] = ctx[ctx_depth] & ^CTX_NEWLINE
                    // Not an indent, start of line, drop out of NL context
                }
            }

            add_to_buf(char)

        }
        // Check if we've hit an atom seperator

    }

    return first_token
}

