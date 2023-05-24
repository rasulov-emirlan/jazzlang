package file

import (
	"fmt"
	"os"

	"github.com/rasulov-emirlan/sunjar/src/evaluator"
	"github.com/rasulov-emirlan/sunjar/src/lexer"
	"github.com/rasulov-emirlan/sunjar/src/macros"
	"github.com/rasulov-emirlan/sunjar/src/object"
	"github.com/rasulov-emirlan/sunjar/src/parser"
)

func EvaluateFile(path string, mp macros.MacrosProcessor) error {
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

	buff = mp.Process(buff)

	// first convert it to string and then to []rune, because we do not want to lose encoding
	l := lexer.New([]rune(string(buff)))
	program := parser.New(l)
	if program == nil {
		return fmt.Errorf("could not parse file %s", path)
	}
	if len(program.Errors()) != 0 {
		return fmt.Errorf("could not parse file %s: %v", path, program.Errors())
	}

	env := object.NewEnvironment()
	res := evaluator.Eval(program.ParseProgram(), env)
	if res != nil && res.Type() == object.OBJ_ERROR {
		fmt.Println(res.Inspect())
	}

	return nil
}
