package bin2dec

import (
	"fmt"
	"strconv"
)

func main() {
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
