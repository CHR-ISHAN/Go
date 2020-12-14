package array_test

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	var arr1 = [4]int{1, 2, 3}
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr1[2])
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	arr4 := [3][3]int{{1, 2, 3}, {1, 2, 3}}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
		t.Log("aaaa")
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
	for _, e := range arr3 {
		t.Log(e)
	}
	t.Log(arr4)
}
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	arr3_sec := arr3[:3]
	arr := arr3[0:1]
	t.Log(arr3_sec)
	t.Log(arr)
}

func f(a [3]int) { fmt.Println(a) }
func fp(a *[3]int) {
	fmt.Println(a)
	a[1] = 2
	fmt.Println(a)
}

func TestArr(t *testing.T) {
	//var ar [3]int
	//var a int
	//f(ar)
	//t.Log(&ar, &a)
	//fp(&ar)
	var arr = []int{1, 2, 3, 4}
	t.Log(sum(arr))
}

func sum(arr []int) int {
	a := 0
	for _, v := range arr {
		a += v
	}
	return a
}
