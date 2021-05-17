package repl

import (
	"bufio"
	"fmt"
	"interpreter_in_go/lexer"
	"interpreter_in_go/token"
	"io"
)

const PROMPT = ">> "

//Start the repl
func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		//wait until type in
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lxr := lexer.NewLexer(line)
		for tkn := lxr.NextToken(); tkn.Type != token.EOF; tkn = lxr.NextToken() {
			fmt.Fprintf(out, "%+v\n", *tkn)
		}
	}
}
