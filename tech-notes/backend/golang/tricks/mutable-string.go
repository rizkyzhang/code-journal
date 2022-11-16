package main

import (
	"fmt"
)

func main() {
	buf := []rune("golang")
	buf[0] = 'G'
	s := string(buf)
	fmt.Println(s) // "Golang"
}
