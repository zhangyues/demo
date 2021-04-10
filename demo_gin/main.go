package main

import "fmt"

func main() {
	one(1, callback)
}

type fb func(int642 int)

type i interface {
	send(s int)
}

// send 发送信息
func (f fb) send(i int) {
	f(i)
}

//需要传递函数
func callback(i int) {
	fmt.Println("i am callBack")
	fmt.Println(i)
}

func two(i int, is i) () {
	is.send(1)
}

func one(int642 int, fc func(int)) {
	two(int642, fb(fc))
}
