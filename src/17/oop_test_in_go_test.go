package _7

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p *Pet) Speak()  {
	fmt.Println("...")
}

type Dog struct {
	p *Pet
}

func (d * Dog) Speak()  {
	//d.p.Speak()
	fmt.Println("wang,wang,wang")
}

type Cat struct {
	Pet
}

func (d * Cat) Speak()  {
	//d.p.Speak()
	fmt.Println("miao,miao,miao")
}

func TestDogSpeak(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	//var Pet cat :=new (Cat)
	//cat.Speak()
}