package main

import (
	"fmt"
	"runtime"
)

func showNumber(num int)  {
	fmt.Println(num)
}

func main() {
	it := 10

	for i:=0; i <= it; i++ {
		go showNumber(i)
	}
	runtime.GOMAXPROCS(2)
	runtime.Gosched()
	fmt.Println("GoodBye")
}
