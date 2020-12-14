package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 3, 5}
	t.Log(a == b)
	t.Log(b == c)
	t.Log(len(a))
}

func TestArray(t *testing.T) {
	t.Log(1 &^ 0) //右边为0则不变，为1则清零
	t.Log(3 &^ 1)
	t.Log(2 &^ 2)
}
