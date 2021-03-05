package _struct

import (
	"fmt"
	"strings"
	"testing"
)

type Employee struct {
	Id   int32
	Name string
	Age  int32
}

type Person struct {
	firstName string
	last      string
}

type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b int
	c int32
	int
	innerS
}

type TwoInt struct {
	a int
	b int
}

type Log struct {
	msg string
}
type Customer struct {
	Name string
	log  *Log
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func TestMagic(t *testing.T) {
	//v := new(Voodoo)
	//v.Magic()
	//v.MoreMagic()

}

func TestCustomer(t *testing.T) {
	c := new(Customer)
	c.Name = "Chris"
	c.log = new(Log)
	c.log.msg = "111"
	t.Log(c.log.Add("Hello"))
}

func (l *Log) Add(s string) string {
	l.msg += "\n" + s
	return l.msg
}

func TestMethod(t *testing.T) {
	two1 := new(TwoInt)
	two1.a = 1
	two1.b = 2
	t.Log(two1.Add(), two1.b)
	t.Log(two1.AddThem(2), two1.b)
}

func (tn *TwoInt) AddThem(ab int) int {
	tn.b = 4
	return tn.a + tn.b + ab
}

func (tn TwoInt) Add() int {
	tn.b = 3
	return tn.b
}

func TestStructInherit(t *testing.T) {
	outer := new(outerS)
	outer.b = 1
	outer.c = 2
	outer.int = 3
	outer.in1 = 4
	outer.in2 = 5
	t.Log(*outer)

}

func (e Employee) String() string {
	return fmt.Sprintf("Id : %d-Name: %s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStringStruct(t *testing.T) {
	e := Employee{0, "Bob", 0}
	t.Log(e)
	t.Log(e.String())
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.last = strings.ToUpper(p.last)
}

func TestStruct(t *testing.T) {
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.last = "Han"
	t.Log(pers1)
	upPerson(&pers1)
	t.Log(pers1)
	pers2 := new(Person)
	pers2.firstName = "Chris"
	(*pers2).last = "han"
	t.Log(pers2)
	t.Log(*pers2)
	test := &Employee{}
	test.Age = 0
	test.Name = "Chris"
	t.Log(test.String())
}
