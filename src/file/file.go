package file

import (
	"fmt"
	"os"

	"github.com/rasulov-emirlan/jazzlang/src/evaluator"
	"github.com/rasulov-emirlan/jazzlang/src/lexer"
	"github.com/rasulov-emirlan/jazzlang/src/object"
	"github.com/rasulov-emirlan/jazzlang/src/parser"
)

func EvaluateFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("'%s' is a directory", path)
	}

	buff := make([]byte, info.Size())
	_, err = f.Read(buff)
	if err != nil {
		return err
	}

	l := lexer.New(string(buff))
	program := parser.New(l)
	env := object.NewEnvironment()
	evaluator.Eval(program.ParseProgram(), env)

	return nil
}
