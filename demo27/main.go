package main

import (
	"fmt"
	"time"
)
//计时程序1
func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 OK"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
		//当接收需求时间比收到时间短，会超时
	case <-time.After(1*time.Second):
		fmt.Println("time out c1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "c2 OK"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
		//发送时间比需求时间短，不会超时
	case <-time.After(4*time.Second):
		fmt.Println("time out c2")
	}
}
