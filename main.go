package main

import (
	"flag"
	"fmt"
)

func main() {
	str := flag.String("str", "", "")
	flag.Parse()
	print(*str)
	fmt.Println("Hello, World!")
}
