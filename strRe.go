package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("Hello", "l", "z", 2))
	fmt.Println(strings.Replace("Wednesday", "e", "xx", -1))
	fmt.Println(strings.Replace("thapthim Santhi", "th", "999", -1))
	fmt.Println(strings.Replace("2521233", "2", "8", 2))
}
