package evaluator

import (
	"testing"

	"github.com/rasulov-emirlan/jazzlang/src/lexer"
	"github.com/rasulov-emirlan/jazzlang/src/object"
	"github.com/rasulov-emirlan/jazzlang/src/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected int64
	}{
		{"simple int", "5", 5},
		{"bigger int", "10", 10},
		{"negative int", "-5", -5},
		{"bigger negative int", "-10", -10},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			evaluated := testEval(tt.input)
			testIntegerObject(t, evaluated, tt.expected)
		})
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}
