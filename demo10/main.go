package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var balance int		// 刚开始的总金额
var transactionNo int	// 最后剩下的金额

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	tranChan := make(chan bool)
	balance = 1000
	transactionNo = 0
	fmt.Println("Starting balance: $",balance)

	wg.Add(1)
	for i := 0; i < 100; i++ {
		go func(ii int, trChan chan bool) {
			//随机取0-25以内的金额
			transactionAmount := rand.Intn(25)
			//用 总金额减去0-25随机金额
			transaction(transactionAmount)
			if ii == 99 {
				trChan <- true
			}
		}(i,tranChan)
	}

	go transaction(0)
	select {

	case <- tranChan:
		fmt.Println("Transactions finished")
		wg.Done()

	}
	wg.Wait()
	close(tranChan)
	fmt.Println("Final balance: $",balance)
}

func transaction(amt int) bool {

	approved := false
	if (balance-amt) < 0 {
		approved = false
	}else {
		approved = true
		balance = balance - amt
	}

	approvedText := "declined"
	if approved == true {
		approvedText = "approved"
	}else {

	}
	transactionNo = transactionNo + 1
	fmt.Println(transactionNo,"Transaction for $",amt,approvedText)
	fmt.Println("\tRemaining balance $",balance)
	return approved
}

