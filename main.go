package main

import (
	"fmt"
)

//func  main()  {
//	fmt.Println("Hello World")
//	os.Exit(1)
//}

//func fib(num int) int {
//	if num < 2 {
//		return 1
//	}
//	return fib(num-1) + fib(num-2)
//}
//
//func main() {
//	for i := 0; i < 2; i++ {
//		nums := fib(i)
//		fmt.Println(nums)
//	}
//	fmt.Println(a)
//}
//func main() {
//	var goos int = runtime.NumCPU()
//	fmt.Printf("The operating system is: %d\n", goos)
//	path := os.Getenv("PATH")
//	fmt.Printf("Path is %s\n", path)
//}

//func main()  {
//	a,b,c := 1,2,3
//	a,b,c = c,b,a
//	fmt.Println(a,b,c)
//}

//var a = "G"
//func main() {
//	n()
//	m()
//	n()
//}
//func n() { print(a) }
//func m() {
//	a := "O"
//	print(a)
//}
//
//func main()  {
//	for i := 0;i<10;i++ {
//		a := rand.Intn(10)
//		fmt.Println(&a)
//	}
//	fmt.Println(time.Now())
//	i := 0
//	switch i {
//	case 0:
//		main()
//	case 1:
//
//	}
//}

//func main() {
//	for i := 0; i < 25; i++ {
//		for j := 0; j < i; j++ {
//			fmt.Print("G")
//		}
//		if i > 10 {
//			break
//		}
//		fmt.Print("\n")
//	}
//}
// this function changes reply:
//func Multiply(a, b int, reply *int) {
//	*reply = a * b
//}
//func min(a ...int) int {
//	if len(a) == 0 {
//		return 0
//	}
//	min := a[0]
//	for _, v := range a {
//		if v < min {
//			min = v
//		}
//	}
//	return min
//}
////func f() {
////	for i := 0; i < 5; i++ {
////		defer fmt.Printf("%d ", i)
////	}
//}
//func func1(s string) (n int, err error) {
//	defer func() {
//		log.Printf("func1(%q) = %d, %v", s, n, err)
//	}()
//	return 7, io.EOF
//}
//
//func fib(n int) (res int) {
//	if n <= 1 {
//		res = 1
//	} else {
//		res = fib(n-1) + fib(n-2)
//	}
//	return res
//}
//func print(n int) (res int) {
//	if n == 0 {
//		return res
//	} else {
//		fmt.Println(n)
//		return print(n - 1)
//	}
//}
//
//func fc() {
//	for i := 0; i < 4; i++ {
//		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
//		g(i)
//		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
//	}
//}
//
//func twoSum(nums []int, target int) []int {
//	for i, v := range nums {
//		for k := i + 1; k < len(nums); k++ {
//			if target-v == nums[k] {
//				return []int{i, k}
//			}
//		}
//	}
//	return []int{}
//}
//
//const (
//	Monday = iota + 1
//	Tues
//	Sec
//)
//const (
//	Open = 1 << iota
//	Close
//)
//
//var a int

func main() {
	//n := 0
	//fmt.Println("Multiply",n)
	//reply := &n
	//Multiply(10, 5, reply)
	//fmt.Println("Multiply", *reply,n) // Multiply: 50
	//x := min(1,2,3,4)
	//fmt.Println(x)
	//f()
	//func1("Hello")
	//fmt.Println(fib(3))
	//print(10)
	//fc()
	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(twoSum(arr, 10))
	//fmt.Println("%d",Tues)
	//fmt.Println(twoSum([]int{1, 2, 3, 4}, 5))
	//f()
	//arr := []float64{7.0, 8.5, 9.1}
	//fmt.Println(Sum(&arr))
	arr := []int{1, 2, 3, 4}
	fmt.Println(sum(arr))
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
func Sum(a *[]float64) (sum float64) {
	for _, v := range *a {
		sum = sum + v
	}
	return
}

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s = s + v
	}
	return s
}
