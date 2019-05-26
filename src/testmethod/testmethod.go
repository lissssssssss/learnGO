package main 
import (
	"fmt"
	"reflect"
)

type S struct{
	name string
}

type T struct{
	S
}

type i interface {
	sVal()
	tVal()
}

func (s S) sVal() {
	fmt.Println("name", s.name, "|")
}
func (s *S) tVal() {
	s.name = "aaa"
	fmt.Println("name", s.name, "|")
}
func (T) tVal() {}
func (T) sVal() {}

func (*S) sPtr() {}
func (*T) tPtr() {}

func MethodSet(a i) {
	t := reflect.TypeOf(a)
	n := t.NumMethod()
	fmt.Println(n)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	var s S
	s.tVal()
	s.sVal()
	fmt.Println(reflect.TypeOf(s).NumMethod())
	//MethodSet(s)
	fmt.Println("*******************")
	MethodSet(&s)
}
