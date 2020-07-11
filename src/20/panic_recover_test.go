package _0

import (
	"errors"
	"fmt"
	"testing"
)

/**
recover得到的异常是panic中创建的
 */

var customError error = errors.New("this is a test error")

func TestRecoverAndDeferAndPanic(t *testing.T) {

	defer func() {
		if err := recover();err ==customError{
			fmt.Println("recover from error")
		}
	}()

	fmt.Println("start test")

	panic(customError)
}