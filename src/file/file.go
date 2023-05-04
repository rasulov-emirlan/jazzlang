package file

import (
	"fmt"
	"log"
	"os"

	"github.com/rasulov-emirlan/sunjar/src/evaluator"
	"github.com/rasulov-emirlan/sunjar/src/lexer"
	"github.com/rasulov-emirlan/sunjar/src/object"
	"github.com/rasulov-emirlan/sunjar/src/parser"
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
	if len(program.Errors()) != 0 {
		log.Println("Parser:", program.Errors())
	}
	env := object.NewEnvironment()
	res := evaluator.Eval(program.ParseProgram(), env)
	if res.Type() == object.OBJ_ERROR {
		log.Println("Evaluator:", res.Inspect())
	}
	return nil
}
