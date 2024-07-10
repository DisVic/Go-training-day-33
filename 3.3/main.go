package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(number uint) uint {
		newNumber := ""
		stringNumber := strconv.FormatUint(uint64(number), 10)
		for i := 0; i < len(stringNumber); i++ {
			if stringNumber[i]%2 == 0 && stringNumber[i] != 48 {
				newNumber += string(stringNumber[i])
			}
		}
		number2, _ := strconv.ParseUint(newNumber, 10, 32)
		if number2 == 0 {
			return 100
		}
		return uint(number2)
	}
	fmt.Printf("%d", fn(200))
}
