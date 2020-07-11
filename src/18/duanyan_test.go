package _8

import (
	"fmt"
	"testing"
)

type AInterface interface {
}

func doSomething(p interface{}) {
	if i, ok := p.(string); ok {
		fmt.Println("String", i)
		return
	}

	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	fmt.Println("other")
}

func TestDuanyan(t *testing.T) {
	doSomething("inter")
	doSomething(1)
}


/**
go中的最佳实践：
1，倾向于使用小的接口定义，很多接口只包含一个方法
2，较大的接口的定义，可以有多个小接口定义组合而成
3，只依赖必要功能的最小接口
 */