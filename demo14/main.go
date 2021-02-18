package main

import "C"
import (
	"fmt"
	"sync"
	"time"
)

//func MyGoFunction(num C.int) int {
//	squared := num * num
//	fmt.Println(num,"squared is",squared)
//	return squared
//}
var (
	mutex sync.Mutex
	wg sync.WaitGroup
)
func hello1(s string)  {
	mutex.Lock()
	for _,i := range s {
		fmt.Printf("%c\n",i)
		time.Sleep(time.Millisecond*2)
	}
	mutex.Unlock()
	wg.Done()
}

func hello2(s string)  {
	mutex.Lock()
	for _,i := range s {
		fmt.Printf("%c\n",i)
		time.Sleep(time.Millisecond*2)
	}
	mutex.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	//v := C.CString("Don't Forget My Memory Is Not Visible To Go!")
	//x := C.string_length(v)
	//fmt.Println("A C function has determined your stringis",x,"characters in length")
	go hello1("world")
	go hello2("hello")
	wg.Wait()

}
