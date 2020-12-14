package loop

import (
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		n++
		t.Log(n)
	}
}

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log(a)
	}
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("aaa")
		case 3, 4:
			t.Log("bbb")
		default:
			t.Log("cccc")
		}
	}
}
