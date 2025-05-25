package main

import (
	"fmt"
	"example.com/sample/lib"
)

func main() {
	str := "!Hello World!"
	str = lib.TrimExclamation(str)
	fmt.Println(str)
}
