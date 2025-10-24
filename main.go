package main

import (
	"fmt"

	"github.com/ayiezac78/golang-walkthrough/helloworld"
	"github.com/ayiezac78/golang-walkthrough/integers"
	"github.com/ayiezac78/golang-walkthrough/iteration"
)

func main() {
	fmt.Println(helloworld.Hello("Ayie", "French"))
	fmt.Println(integers.Add(2, 2))
	fmt.Println(iteration.Repeat("a"))
}
