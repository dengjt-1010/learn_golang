package _13

import (
	"fmt"
	"testing"
	"time"
)


/**
函数式编程
 */

/**
1，函数可以有多个返回值
2，所有参数都时值传递
3，函数可以作为变量的值
4，函数可以作为参数和返回值
*/

func returnMultiValues(input int) (ret int) {
	return input * input
}

/**
简直牛逼，再也不担心在方法周围做增强的。
*/
func timspant(innter func(op int) int) func(op int) int {

	return func(op int) int {
		start := time.Now()
		ret := innter(op)
		fmt.Print("time spent", time.Since(start))
		return ret
	}
}

func TestReturnMultiValues(t *testing.T) {
	newfun := timspant(returnMultiValues)
	ret := newfun(10)
	t.Log(ret)
}


/**
可变参数
 */

func sum(params ... int) int  {
	var sum int
	for _,value :=range params{
		sum+=value
	}
	return sum
}

func TestVariableParam(t *testing.T)  {
	t.Log(sum(100,200,300))
}

/**
defer类似于java中的finally
 */

func TestDefer(t *testing.T)  {
	defer func() {
		t.Log("clear some resource")
	}()

	t.Log("start")

	panic("err")
}