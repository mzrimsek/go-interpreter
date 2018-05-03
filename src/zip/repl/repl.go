// Package repl : The interactive language shell to work with Zip code
package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"zip/evaluator"
	"zip/lexer"
	"zip/object"
	"zip/parser"
)

// PROMPT : The characters representing the start of each input line in the REPL
const PROMPT = ">> "

// Start : Starts an interactive shell session for a user to input Zip code into
func Start(in io.Reader, out io.Writer) {
	evaluator.ShouldPrint = true

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! Something went terribly wrong here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
