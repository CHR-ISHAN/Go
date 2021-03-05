package _interface

import (
	"fmt"
	"testing"
)

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func TestInterfaces(t *testing.T) {
	//sq1 := new(Square)
	//sq1.side = 5
	//var areaIntf Shaper
	//areaIntf = sq1
	//t.Log(areaIntf.Area())
	r := Rectangle{width: 5, length: 3}
	q := &Square{side: 5}
	shapes := []Shaper{r, q}
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

type car struct {
	make  string
	model string
	price float32
}

func (s *stockPosition) getValue() float32 {
	return s.count * s.sharePrice
}

func (c *car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(T valuable) {
	fmt.Printf("Value of the asset is %f\n", T.getValue())
}

func TestInter(t *testing.T) {
	var o valuable = &stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = &car{"BMW", "M3", 66500}
	showValue(o)
}

type List []int

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func (l List) Len() int {
	return len(l)
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func TestSetInterface(t *testing.T) {

	var lst List
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {

		fmt.Printf("- plst is long enough\n")
	}
}
