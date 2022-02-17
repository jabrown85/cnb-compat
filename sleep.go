package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if args := len(os.Args); args != 1 {
		fmt.Println("Args are not allowed!")
		os.Exit(1)
	}
	fmt.Println("Sleeping forever")
	time.Sleep(24 * time.Hour)
	fmt.Println("Completed")
}
