package main

import "fmt"

var comm = make(chan bool)
var done = make(chan bool)

func prodecuer()  {
	for i:=0; i<10; i++ {
		comm <- true
	}
	done <- true
}
func consumer()  {
	for {
		communication := <-comm
		fmt.Println("Communication from producerreceived!",communication)
	}
}

func main() {
	go prodecuer()
	go consumer()
	<- done
	fmt.Println("All done!")
}
