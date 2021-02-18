package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var initialString string
var finalString string

var stringLength int

func addToFinalStack(letterChannel chan string, wg *sync.WaitGroup)  {
	letter := <- letterChannel
	finalString += letter
	wg.Done()
}

func captialize(letterChannel chan string, currentString string,wg *sync.WaitGroup)  {
	thisLetter := strings.ToUpper(currentString)
	wg.Done()
	letterChannel <- thisLetter
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	initialString = "Four score and seven years ago our fathersbrought " +
		"forth on this continent, a new nation, conceived inLiberty, " +
		"and dedicated to the proposition that all men arecreated equal."
	initialByte := []byte(initialString)
	var letterChannel = make(chan string)
	stringLength = len(initialByte)
	for i:=0; i<stringLength; i++ {
		wg.Add(2)
		go captialize(letterChannel, string(initialByte[i]), &wg)
		go addToFinalStack(letterChannel, &wg)
		wg.Wait()
	}
	fmt.Println(finalString)
}
