package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if args := len(os.Args); args != 1 {
		fmt.Println("Args are not allowed!")
		fmt.Println("I will exit in 5 minutes with status code 1")
		time.Sleep(5 * time.Minute)
		os.Exit(1)
	}
	fmt.Println("Sleeping forever")
	time.Sleep(24 * time.Hour)
	fmt.Println("Completed")
}
