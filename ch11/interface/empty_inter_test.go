package _interface

import (
	"fmt"
	"reflect"
	"testing"
)

var i = 5

var str = "ABC"

type Any interface {
}

type Element interface {
}

type Vector struct {
	a []Element
}

func (p *Vector) At(i int) Element {
	return p.a[i]
}

func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

type Person struct {
	name string
	age  int
}

type IDuck interface {
	Quack()
	Walk()
}

type Bird struct {
}

type duck struct {
}

func (b *duck) Quack() {
	fmt.Println("I am quacking!")
}

func (b *duck) Walk() {
	fmt.Println("I am walking!")
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!11111")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking!11111")
}

func DuckDance(duck IDuck) {
	duck.Quack()
	duck.Walk()
}

func TestDuck(t *testing.T) {
	b := new(Bird)
	DuckDance(b)
}

type MyInt int

var m MyInt = 5

func TestEmpty(t *testing.T) {
	//var val Any
	//val = 5
	//t.Log(val)
	//val = str
	//t.Log(val)
	//pers1 := new(Person)
	//pers1.age = 10
	//pers1.name = "Chris"
	//val = pers1
	//t.Log(val)
	//
	//switch val.(type) {
	//case int:
	//	t.Log("int")
	//case string:
	//	fmt.Printf("Type string %T\n", t)
	//case bool:
	//	fmt.Printf("Type boolean %T\n", t)
	//case *Person:
	//	fmt.Printf("Type pointer to Person %T\n", t)
	//default:
	//	fmt.Printf("Unexpected type %T", t)
	//}
	a := 5
	v := reflect.ValueOf(m)
	t.Log(v, reflect.TypeOf(a))
}
