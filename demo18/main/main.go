package main

import (
	"fmt"
	"sync"
	"time"
)

type TimeStruct struct {
	totalChange int
	currentTime time.Time
	rw sync.RWMutex
}
var TimeElement TimeStruct
func updateTime()  {
	TimeElement.rw.Lock()
	defer TimeElement.rw.Unlock()
	TimeElement.currentTime = time.Now()
	TimeElement.totalChange++
}

func main() {
	var wg sync.WaitGroup
	TimeElement.currentTime = time.Now()
	TimeElement.totalChange = 0
	timer := time.NewTicker(1 * time.Second)
	writerTimer := time.NewTicker(10 * time.Second)
	endTimer := make(chan bool)

	wg.Add(1)
	go func() {
		for {
			select {
				case <-timer.C:
					fmt.Println(TimeElement.totalChange, TimeElement.currentTime.String())
				case <-writerTimer.C:
					updateTime()
				case <-endTimer:
					timer.Stop()
					return
			}
		}
	}()

	wg.Wait()
	fmt.Println(TimeElement.currentTime.String())
}
