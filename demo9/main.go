package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
/*	current := 0
	it := 100
	var wg sync.WaitGroup
	for i:=0; i<it; i++ {
		wg.Add(1)

		go func() {
			current++
			fmt.Println(current)
			wg.Done()
		}()
		wg.Wait()
	}*/
	runtime.GOMAXPROCS(2)
	current := 0
	it := 100
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(it)
	for i:=0; i<it; i++ {
		go func() {
			mutex.Lock()
			fmt.Println(current)
			current++
			mutex.Unlock()
			fmt.Println(current)
			wg.Done()
		}()
	}
	wg.Wait()
}
