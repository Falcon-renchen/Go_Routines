package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
func A()  {
	for i:=1; i<=10; i++ {
		fmt.Println("A",i)
	}
	wg.Done()
}
func B()  {
	for i:=1; i<=10; i++ {
		fmt.Println("B",i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(6)	//占用6个cpu核心
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
}