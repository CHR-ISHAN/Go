package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "中华"
	c := []rune(s)
	t.Log(s, c)
}

func TestStringFunc(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)
	t.Log(strings.Join(parts, "."))
	i := 1234
	t.Log(strconv.Itoa(i) + "aaa")
}

func Return() string {
	return "Hello"
}
