package main

import (
	"fmt"
	"os"
)

/**
	应用程序入口：
		1，必须是main包，
		2，必须是main方法
		3，文件名不一定是main.go
 */

/**
 不同与java，package与目录的路径可以不一致。
 */
func main () {
	fmt.Println(len(os.Args))//可以通过启动命令 go run helloworld.go 666 来读取"666" 参数
	fmt.Println("hello world")
	os.Exit(123) //程序退出的时候可以返回状态码
}
