package main

import (
	"fmt"

	"github.com/hogehoge/gomono/foo"

	// Not Support rename import
	//teta "github.com/hogehoge/gomono/foo/glbl"
	// Not Support blank import
	//import _ "net/http/pprof"
	// Not Support dot import
	//import . "fmt"
	"github.com/hogehoge/gomono/foo/glbl"
)

func main() {
	a := MainUtilMethod("ok", "ng")

	fmt.Println(a)

	queue := foo.NewQueue[int](0)

	b := &foo.Queue[int]{}

	var c int

	if b.Empty() {
		c = queue.Len()
	}

	fmt.Println(c)

	fmt.Println(tete.A)

	return
}
