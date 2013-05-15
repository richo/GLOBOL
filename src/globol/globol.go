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
        /* err error */
        /* part []byte */
        /* prefix bool */
    )

    buf := make([]byte, 1024)

    file, _ = os.Open(os.Args[1])
    for {
        n, err := file.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }

        // Iterate over bytes read
        i = 0
        for i < n {
            char = buf[i]
            if lexer.IsAtomSeperator(char) {
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
