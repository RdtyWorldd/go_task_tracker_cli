package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Arguments incorect. \n Use -h flag to get help information")
	}

}
