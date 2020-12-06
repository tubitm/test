package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("Hello", "l", "z", 2))
	fmt.Println(strings.Replace("Wednesday", "e", "xx", -1))
	fmt.Println(strings.Replace("thapthim Santhi", "Th", "999", -1))
}
