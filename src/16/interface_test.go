package _6

import "testing"

type Programmer interface {
	writeHelloWorld() string
}

type GoProgrammer struct {

}

func (p *GoProgrammer) writeHelloWorld() string  {
	return "go: hello world"
}

func TestImplementFunc(t *testing.T) {
	var programmer Programmer
	programmer = new(GoProgrammer)
	t.Log(programmer.writeHelloWorld())
}

/**
1，接口为非入侵性，实现不依赖接口的定义
2，接口的定义可以包含在使用者的包内
 */