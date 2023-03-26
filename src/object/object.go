package object

import (
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/rasulov-emirlan/sunjar/src/ast"
)

type ObjectType string

const (
	OBJ_INTEGER      = "INTEGER"
	OBJ_STRING       = "STRING"
	OBJ_BOOLEAN      = "BOOLEAN"
	OBJ_NULL         = "NULL"
	OBJ_RETURN_VALUE = "RETURN_VALUE"
	OBJ_FUNCTION     = "FUNCTION"
	OBJ_ERROR        = "ERROR"
	OBJ_BUILTIN      = "BUILTIN"
	OBJ_ARRAY        = "ARRAY"
	OBJ_HASH         = "HASH"
	OBJ_LOOP         = "LOOP"
)

type (
	Object interface {
		Type() ObjectType

		// Returns a string representation of the object.
		// Most of the time, this will be the literal value of the object.
		Inspect() string
	}

	Hashable interface {
		HashKey() HashKey
	}
)

type Integer struct {
	Value int64
}

var _ Object = (*Integer)(nil)

func (i *Integer) Type() ObjectType { return OBJ_INTEGER }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

type String struct {
	Value string
}

var _ Object = (*String)(nil)
var _ Hashable = (*String)(nil)

func (s *String) Type() ObjectType { return OBJ_STRING }
func (s *String) Inspect() string  { return s.Value }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type Boolean struct {
	Value bool
}

var _ Object = (*Boolean)(nil)
var _ Hashable = (*Boolean)(nil)

func (b *Boolean) Type() ObjectType { return OBJ_BOOLEAN }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) HashKey() HashKey {
	value := 0
	if b.Value {
		value = 1
	}
	return HashKey{Type: b.Type(), Value: uint64(value)}
}

type Null struct{}

var _ Object = (*Null)(nil)

func (n *Null) Type() ObjectType { return OBJ_NULL }
func (n *Null) Inspect() string  { return "null" }

type ReturnValue struct {
	Value Object
}

var _ Object = (*ReturnValue)(nil)

func (rv *ReturnValue) Type() ObjectType { return OBJ_RETURN_VALUE }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

var _ Object = (*Function)(nil)

func (f *Function) Type() ObjectType { return OBJ_FUNCTION }
func (f *Function) Inspect() string {
	var out string

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out += fmt.Sprintf("fn(%s) {\n", strings.Join(params, ", "))
	out += f.Body.String()
	out += "}"

	return out
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

var _ Object = (*Builtin)(nil)

func (b *Builtin) Type() ObjectType { return OBJ_BUILTIN }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

var _ Object = (*Array)(nil)

func (ao *Array) Type() ObjectType { return OBJ_ARRAY }

func (ao *Array) Inspect() string {
	var out string

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out += fmt.Sprintf("[%s]", strings.Join(elements, ", "))

	return out
}

// HASH MAPS
type (
	HashKey struct {
		Type  ObjectType
		Value uint64
	}

	HashPair struct {
		Key   Object
		Value Object
	}

	Hash struct {
		Pairs map[HashKey]HashPair
	}
)

var _ Object = (*Hash)(nil)

func (h *Hash) Type() ObjectType { return OBJ_HASH }
func (h *Hash) Inspect() string {
	var out string

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out += fmt.Sprintf("{%s}", strings.Join(pairs, ", "))

	return out
}

type Loop struct {
	Condition ast.Expression
	Body      *ast.BlockStatement
	Env       *Environment
}

var _ Object = (*Loop)(nil)

func (l *Loop) Type() ObjectType { return OBJ_LOOP }
func (l *Loop) Inspect() string {
	var out string

	out += fmt.Sprintf("loop(%s) {\n", l.Condition.String())
	out += l.Body.String()
	out += "}"

	return out
}

type Error struct {
	Message string
}

var _ Object = (*Error)(nil)

func (e *Error) Type() ObjectType { return OBJ_ERROR }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
