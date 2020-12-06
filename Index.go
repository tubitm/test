package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("Hello World", "hello"))
	fmt.Println(strings.Index("sunday", "day"))
	fmt.Println(strings.Index("John H. Watson", "watson"))
	fmt.Println(strings.Index("John H. Watson", "."))
}
