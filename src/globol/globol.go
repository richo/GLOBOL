/*

Main entry point to globol

*/

package main

import (
 	 "os"
 	 "fmt"
 	 "globol/lexer"
 	 "globol/parser"
)

func usage() {
 	 fmt.Println("Usage:")
 	 fmt.Println("  globol lex <file> 	  	 : Lex file and print to stdout")
 	 fmt.Println("  globol parse <file> 	   : Parse file and print to stdout")
 	 os.Exit(1)
}

func lex(file *os.File) {
 	 var (
 	  	 token_list *lexer.Token
 	 )

 	 token_list = lexer.Lex(file)

 	 current_token := token_list

 	 /* first_token := token_list */

 	 for {
 	  	 fmt.Println(current_token.Type, string(current_token.Content))

 	  	 current_token = current_token.Next
 	  	 if (current_token == nil) {
 	  	  	 break
 	  	 }
 	 }
}

func parse(file *os.File) {
 	 var (
 	  	 token_list *lexer.Token
 	  	 ast *parser.AST
 	 )
 	 token_list = lexer.Lex(file)
 	 ast = parser.Parse(token_list)

 	 fmt.Println(ast)
}

func main() {
 	 var (
 	  	 file *os.File
 	 )

 	 if (os.Args[1] == "lex") {
 	  	 file, _ = os.Open(os.Args[2])
 	  	 lex(file)
 	 } else
 	 if (os.Args[1] == "parse") {
 	  	 file, _ = os.Open(os.Args[2])
 	  	 parse(file)
 	 } else {
 	  	 usage()
 	 }
}
