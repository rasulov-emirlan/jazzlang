package object

import "fmt"

type ObjectType string

const (
	OBJ_INTEGER      = "INTEGER"
	OBJ_BOOLEAN      = "BOOLEAN"
	OBJ_NULL         = "NULL"
	OBJ_RETURN_VALUE = "RETURN_VALUE"
)

type Object interface {
	Type() ObjectType

	// Returns a string representation of the object.
	// Most of the time, this will be the literal value of the object.
	Inspect() string
}

type Integer struct {
	Value int64
}

var _ Object = (*Integer)(nil)

func (i *Integer) Type() ObjectType { return OBJ_INTEGER }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

var _ Object = (*Boolean)(nil)

func (b *Boolean) Type() ObjectType { return OBJ_BOOLEAN }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

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
