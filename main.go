package main

import (
	"fmt"
	"os"
	"strconv"
)

func checknum() {
	var num string
	for {
		fmt.Print("Enter Binary Number:")
		fmt.Scanln(&num)
		output, err := strconv.ParseInt(num, 2, 64)
		if err == nil {
			fmt.Printf("Output: %d", output)
			os.Exit(0)
		}
		if err != nil {
			fmt.Println("invalid input")
		}
	}

}

func main() {
	for {
		fmt.Println("Press 1 to run")
		fmt.Println("Press 2 to exit")
		var input int
		fmt.Scanln(&input)
		if input < 1 {
			fmt.Println("invalid input")
		} else if input > 2 {
			fmt.Println("invalid input")
		} else if input == 1 {
			checknum()
		} else {
			break
		}
	}

}
