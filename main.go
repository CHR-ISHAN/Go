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

type TT float64

func (t TT) String() string {
	return fmt.Sprintln("%v", t)
}

//func main() {
//	//var t TT
//	//t.String()
//	//var m runtime.MemStats
//	//runtime.ReadMemStats(&m)
//	//fmt.Printf("%d Kb\n", m.Alloc/1024)
//}

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

type A interface {
}

type Cat struct {
	name string
	age  int
}

type Person struct {
	name string
	sex  string
}

func test1(a A) {

}

func test2(a interface{}) {

}

type Shaper interface {
	Area() float32
}

type TopologicalGenus interface {
	Rank() int
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Rank() int {
	return 1
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r Rectangle) Rank() int {
	return 2
}

func main() {
	var a = new(ListNode)
	showNode(a)
}

//func twoSum(nums []int, target int) []int {
//	for i, x := range nums {
//		for j := i + 1; j < len(nums); j++ {
//			if x+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, x := range nums {
		if p, ok := hash[target-x]; ok {
			return []int{p, i}
		}
		hash[x] = i
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func showNode(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func insertNode(data int, l *ListNode) {
	tmp := new(ListNode)
	if tmp == nil {
		fmt.Println("err: out of space")
	}
	tmp.Val = data
	tmp.Next = l.Next
	l.Next = tmp
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, carry/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}
