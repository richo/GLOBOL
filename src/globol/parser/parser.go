package parser

import (
 	 "globol/lexer"
)

type AST struct {
 	 Type int
}

func Parse (tokens *lexer.Token) *AST {
 	 var (
 	  	 ast *AST
 	  	 this *lexer.Token
 	 )
 	 this = tokens
 	 for {





 	  	 // Line up the next token
 	  	 if (this.Next == nil) {
 	  	  	 break
 	  	 } else {
 	  	  	 this = this.Next
 	  	 }
 	 }

 	 ast = new(AST)

 	 ast.Type = 1

 	 return ast
}
