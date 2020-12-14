package function_test

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

//var num1 float64
//var num2 int
//var num3 int

func TestReturn(t *testing.T) {
	//num2, num3 = get2and3(2)
	//t.Log(num2, num3)
	//num2, num3 = get2and3_2(3)
	//t.Log(num2, num3)
	//num1, num2, num3 = mult_returnval(3, 2)
	//t.Log(num1, num2, num3)
	//num1, num2, num3 = mult_returnval_2(3, 2)
	//t.Log(num1, num2, num3)
	num1, err := MySqrt_2(3.4123)
	t.Log(num1, err)
	num1, err = MySqrt_2(-1)
	t.Log(num1, err)
}

func TestOutside(t *testing.T) {
	n := 10
	rel := &n
	t.Log(n)
	Multiply(1, 2, rel)
	t.Log(n)
}

func get2and3(input int) (int, int) {
	return 2 * input, 3 * input
}

func get2and3_2(input int) (x2 int, x3 int) { //推荐使用
	x2 = input * 2
	x3 = input * 3
	return
}

func mult_returnval(long int, width int) (add int, sub int, pro int) {
	add = long + width
	sub = long * width
	pro = long - width
	return
}
func mult_returnval_2(long int, width int) (int, int, int) {
	return long + width, long * width, long - width
}

func MySqrt(input float64) (ret float64, err error) {
	if input < 0 {
		ret = math.NaN()
		err = errors.New("error")
	} else {
		ret = math.Sqrt(input)
	}
	return
}

func MySqrt_2(input float64) (float64, error) {
	if input < 0 {
		return math.NaN(), errors.New("error")
	} else {
		return math.Sqrt(input), nil
	}
}
func Multiply(a, b int, n *int) {
	*n = a * b
}

func TestFunc(t *testing.T) {
	t.Log(min(1, 2, 3, 4, 5))
	test := []int{1, 2, 3, 4}
	t.Log(min(test...))
	b()
	//t.Log(fib(10000))
	t.Log(fib(5))
}

func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func fib(num int) (res int, pos int) {
	res = num
	if num <= 2 {
		res = 1
	} else {
		_, v1 := fib(num - 1)
		_, v2 := fib(num - 2)
		pos = v1 + v2
	}
	return
}

func TestCallBack(t *testing.T) {
	//callback(1, Add)
	//t.Log(strings.IndexFunc(" ", unicode.IsSpace))
	//f := func() (int, int) {
	//	sum := 0
	//	for i := 1; i < 1000; i++ {
	//		sum = sum + i
	//	}
	//	return sum, 1
	//}
	//t.Log(f())
	t.Log(f())
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return
}

func TestFunction_return(t *testing.T) {
	p2 := Add2()
	t.Log(p2(3))
	TwoAdder := Adder(2)
	t.Log(TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}
func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
