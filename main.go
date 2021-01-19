package main

import (
	"fmt"
	"strconv"
)

func checknum() {
	var num string
	fmt.Print("Enter Binary Number:")
	fmt.Scanln(&num)
	output, err := strconv.ParseInt(num, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Output %d", output)
}

func main() {
	fmt.Println("Press 1 to run")
	fmt.Println("Press 2 to exit")
	var input int
	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
		fmt.Println("invalid input")
		return
	}
}
