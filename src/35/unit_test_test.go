package _35

import (
	"fmt"
	"testing"
)

func TestSqurare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := squre(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, expected is %d, the actual is %d", inputs[i], expected[i], ret)
		}
	}
}

func squre(input int) int {
	return input * input+1
}


/**
Fail ,Error : 该测试失败，该测试继续，其他测试继续
FailNow,Fatal: 该 测试失败，该测试种植，其他测试继续执行
 */

func TestError(t *testing.T) {
	fmt.Println("start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFailNow(t *testing.T) {
	fmt.Println("start")
	t.FailNow()
	fmt.Println("End")
}

/**
 类似与java中的unit-test的断言assert，golang中也有类似的实现 http://github.com/stretchr/testify
 */