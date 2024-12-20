package main

import
(
	"fmt"
	"gopkg.in/gcfg.v1"
)

type Configuration struct {
	Revisions struct {
		Count int
		Revisionsuffix string
		Lockfiles bool
	}
	Logs struct {
		Rotatelength int
	}
	Alarms struct {
		Emails string
	}
}

func main() {
	configFile := Configuration{}
	err := gcfg.ReadFileInto(&configFile, "example.ini")
	if err != nil {
		fmt.Println("Error",err)
	}
	fmt.Println("Rotation duration:",configFile.Logs.Rotatelength)
}