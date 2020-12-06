package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("ComPro", "o"))
	fmt.Println(strings.Count("Hello World", "l"))
	fmt.Println(strings.Count("Hello World", ""))
	fmt.Println(strings.Count("SunDay", "d"))
}
