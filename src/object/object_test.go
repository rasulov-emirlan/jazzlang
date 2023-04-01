package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringHashKey(t *testing.T) {
	// these tests are to make sure similar keys are hashed to the same value
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	assert.Equal(t, hello1.HashKey(), hello2.HashKey())
	assert.Equal(t, diff1.HashKey(), diff2.HashKey())
	assert.NotEqual(t, hello1.HashKey(), diff1.HashKey())
}

func TestNumbersHashKey(t *testing.T) {
	// make sure that integers and floats are hashed differently
	hello1 := &Integer{Value: 1}
	hello2 := &Integer{Value: 1}
	diff1 := &Integer{Value: 2}
	diff2 := &Integer{Value: 2}
	f1 := &Float{Value: 1.0}
	f2 := &Float{Value: 1.0}

	assert.Equal(t, hello1.HashKey(), hello2.HashKey())
	assert.Equal(t, diff1.HashKey(), diff2.HashKey())
	assert.NotEqual(t, hello1.HashKey(), diff1.HashKey())
	assert.Equal(t, f1.HashKey(), f2.HashKey())
	assert.NotEqual(t, hello1.HashKey(), f1.HashKey())
}
