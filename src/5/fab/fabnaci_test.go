package fab

import (
	"testing"
)
//
//func TestFabonachi(t *testing.T) {
//	var a = 1
//	var b = 1
//	for i := 0; i < 10; i++ {
//		fmt.Println("", b)
//		temp := a
//		a = b
//		b = a + temp
//	}
//}

func Test_fabnaci(t *testing.T) {
	a := 1
	b := 2
	temp := a
	a = b
	b = temp

	t.Log(a,b)
}
