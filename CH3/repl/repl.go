package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// https://fsymbols.com/generators/tarty/
const ERROR_MSG = `

███████╗██████╗░██████╗░░█████╗░██████╗░
██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔══██╗
█████╗░░██████╔╝██████╔╝██║░░██║██████╔╝
██╔══╝░░██╔══██╗██╔══██╗██║░░██║██╔══██╗
███████╗██║░░██║██║░░██║╚█████╔╝██║░░██║
╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝░╚════╝░╚═╝░░╚═╝
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, ERROR_MSG)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

/*
let x = 1 * 2 * 3 * 4 * 5
x * y / 2 + 3 * 8 - 123
true == false
*/

/*
let x 12 * 3
*/

/*
5
10
999
*/

/*
5*5+10
(10 +2) *30 == 300 + 20 * 3
500/2 != 250
*/

/*
if (1 > 2) { 10 } else { 20 }
if ((10 +2) *30 == 300 + 20 * 3) { 1111 } else { 2222 }
if ((10 +2) *30 == 300 + 20 * 31) { 1111 } else { 2222 }
*/
