package typetest

import "testing"

//TODO golang中不支持隐式类型转换,如下会有编译错误
//func TestExchange(t *testing.T) {
//	var a int = 1
//	var b int32 = 1
//	a = b
//	t.Log(a)
//}

// TODO 看来golang中支持强制类型转换
func TestExchange2(t *testing.T) {
	var a int = 1
	var b int32
	b = int32(a)
	t.Log(b)
}

func TestPoint(t *testing.T) {
	a := 1
	aPointer :=&a

	//aPointer = aPointer+1 ,TODO golang中不支持指针运算
	t.Log(a, aPointer)
}