package _5

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

/**
这一节主要看 golang中对象的封装，包括对数据和行为的封装
*/
type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e *Employee) toString1() string {
	fmt.Println("Adress is", unsafe.Pointer(&e.Name))
	return "" + e.Name + e.Id + strconv.Itoa(e.Age)
}

func (e Employee) toString2() string {
	fmt.Println("Adress is", unsafe.Pointer(&e.Name))
	return "" + e.Name + e.Id + strconv.Itoa(e.Age)
}

func TestObject(t *testing.T) {

	e := Employee{"0", "dengjt", 18}
	e1 := Employee{Id: "1", Name: "dengjt", Age: 18}
	e2 := new(Employee)
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Log(&e)

	fmt.Println("Adress is", unsafe.Pointer(&e2.Name))
	t.Log(e2.toString1())
	t.Log(e2.toString2())
}
