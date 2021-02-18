package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello()  {
	fmt.Println("hello 娜扎")
	wg.Done() 	// 技术牌-1
}
func hello2()  {
	fmt.Println("hello 汤姆")
	wg.Done()
}

func main() {

	wg.Add(2)	//技术牌+1

	go hello()
	go hello2()
	fmt.Println("hello main")
	//time.Sleep(time.Millisecond*2000)

	wg.Wait() 	//阻塞   等待技术牌减为0
}
