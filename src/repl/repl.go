package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/rasulov-emirlan/jazzlang/src/evaluator"
	"github.com/rasulov-emirlan/jazzlang/src/lexer"
	"github.com/rasulov-emirlan/jazzlang/src/parser"
)

const PROMPT = "🎵 "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			io.WriteString(out, "Bye!👋\n")
			return
		}

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

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s \n", msg)
	}
}
