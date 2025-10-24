package main

import (
	"fmt"

	// "github.com/ayiezac78/golang-walkthrough/helloworld"
	// "github.com/ayiezac78/golang-walkthrough/integers"
	// "github.com/ayiezac78/golang-walkthrough/iteration"
	"github.com/ayiezac78/golang-walkthrough/sum"
)

func main() {
	// fmt.Println(helloworld.Hello("Ayie", "French"))
	// fmt.Println(integers.Add(2, 2))
	// fmt.Println(iteration.Repeat("a"))
	numbers := []int{1, 2, 3}
	fmt.Println(sum.SumAllTails(numbers, []int{4, 5, 6}))
}
