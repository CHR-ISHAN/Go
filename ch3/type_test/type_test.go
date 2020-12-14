package type_test

import (
	"math"
	"testing"
)

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	c := math.MaxInt64
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(aPtr, a)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
}
