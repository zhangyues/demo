package main

import (
	"fmt"
	"sync"
)

var ch1 chan string
var ch2 chan string
var ch3 chan struct{}
var wg sync.WaitGroup

func main() {
	ch1 = make(chan string)
	ch2 = make(chan string)
	ch3 = make(chan struct{})
	go testChanelCache()
	go testNoCacheChannel()
	for {
		select {
		case o := <-ch1:
			fmt.Println("ch1: ", o)
		case t := <-ch2:
			fmt.Println("ch2: ", t)
		case <-ch3:
			fmt.Println("ch3: " )
		default:
			break
		}
	}

}

// testChanelCache 测试存在缓冲的channel
func testChanelCache() {

	ch1 <- "helloworld"
}

// testNoCacheChannel测试无缓冲channel
func testNoCacheChannel() {

	ch2 <- "a"
	ch2 <- "b"
	ch2 <- "c"
}
