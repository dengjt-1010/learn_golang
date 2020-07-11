package map2_test

import "testing"

//map的value可以是一个函数，轻易的实现工厂模式

func TestMapWithFunValue(t *testing.T) {

	m:=map[int] func(op int)int{}

	m[1] = func(op int) int {
		return op
	}

	m[2]= func(op int) int {

		return op*op
	}
	t.Log(m[1](2))
	t.Log(m[2](2))
}

//golang中是没有set的,可以通过map来实现

func TestSetInGo(t *testing.T) {

	set:=map[int]bool{}

	set[1]=true
	set[3]=true

	if k,ok :=set[1];ok{
		t.Log(k,ok,"is exist")
	}

	if set[3]{
		t.Log("see 3 in set")
	}

	t.Log(len(set))

	delete(set,1)
	t.Log(set)

}