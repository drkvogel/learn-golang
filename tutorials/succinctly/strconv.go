package main

import (
	"fmt"
	"strconv"
)

func main() {
	int1 := 48
	int1_as_string := strconv.Itoa(int1)
	str1 := "My number is "
	str1 = str1 + int1_as_string
	fmt.Println(str1)
}