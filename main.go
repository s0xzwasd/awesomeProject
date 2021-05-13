package main

import "flag"

func main() {
	str := flag.String("str", "", "")
	flag.Parse()
	print(*str)
}
