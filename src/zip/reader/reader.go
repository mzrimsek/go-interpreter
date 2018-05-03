// Package reader : Handles evaluating an inputted Monkey file
package reader

import (
	"bufio"
	"io"
	"zip/evaluator"
	"zip/lexer"
	"zip/object"
	"zip/parser"
)

// Start : Reads and evaluates the inputted file
func Start(in io.Reader, out io.Writer) {
	evaluator.ShouldPrint = false

	scanner := bufio.NewScanner(in)

	var programInput string
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()
		programInput += line
	}

	l := lexer.New(programInput)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		inspected := evaluated.Inspect()
		if inspected != "null" {
			io.WriteString(out, inspected)
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
