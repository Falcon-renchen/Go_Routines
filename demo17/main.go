package main

import (
	"fmt"
	"strings"
)

func shortenString(message string) func() string {
	return func() string {
		messageStirng := strings.Split(message," ")
		wordlen := len(messageStirng)
		if wordlen < 0 {
			return "do nothing"
		} else {
			messageStirng = messageStirng[:(wordlen-1)]
			message = strings.Join(messageStirng, " ")
			return message
		}

	}
}

func main() {
	myString := shortenString("Welcome to me")

	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
}
