package slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0)) //cap会2倍增长
	s0 = append(s0, 1)
	t.Log(s0)
	s2 := make([]int, 3, 5)
	t.Log(s2)
}

func TestSliceShare(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}
	Q2 := year[2:5]
	Q1 := year[1:3]
	t.Log(Q2)
	t.Log(len(Q2), cap(Q2))
	t.Log(Q1)
	year[2] = "AAA"
	t.Log(Q1, Q2)
}

func TestPtr(t *testing.T) {
	//s := "hello"
	//var p *string = &s
	//t.Log(p, *p)
	//*p = "lll"
	//t.Log(*p)
	//item := [...]int{1, 2, 3, 4}
	//for _, v := range item {
	//	v = v * 2
	//}
	//t.Log(item)
	//slice := make([]int, 0, 10)
	//for i := 20; i > 0; i-- {
	//	slice = append(slice, i)
	//	t.Log(len(slice), cap(slice))
	//}
	//t.Log(slice)
	//sort.Ints(slice)
	//t.Log(slice)
	//b := []int{1, 2, 3, 4}
	a := []int{5, 6, 7}
	//a = append(a, b...)
	a = append(a[:2], a[3:]...)
	t.Log(a)
	s1, s2 := sub("hello", 2)
	t.Log(s1, s2)
}

func sub(str string, i int) (c1 string, c2 string) {
	c := []byte(str)
	c1 = string(c[:i])
	c2 = string(c[i:])
	return
}
