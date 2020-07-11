package _22

/**
golang中的这种依赖管理无法解决同一台开发机器上依赖不同版本的库的问题，也就是没有版本的概念
所以，在go 1.5中，vendor目录添加到除了 GOPATH和GOROOT之外的依赖目录查找的方案

1，首先查找当前包vendor目录
2，向上级目录查找，直到找到src下的vendor目录
3，在GOPATH中找
4，在GOROOT目录下找

常用依赖管理工具
godep
glide
dep

 */