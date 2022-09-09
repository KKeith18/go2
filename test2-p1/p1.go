package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func build_slice(x int) []int {
	slice := make([]int, x)
	for i := 0; i <= x-1; i++ {
		slice[i] = rand.Int()
	}
	return slice
}
func add_slice(x []int) int {
	sum := 0
	for i := 0; i <= len(x)-1; i++ {
		sum = x[i] + sum
	}
	return sum
}

// sourced from a tour of go
func summ(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	//send sum to c
	c <- sum
}

func main() {
	//this line copied from go documenation
	firstArg := os.Args[1]
	fmt.Println(firstArg)
	z, err := strconv.Atoi(firstArg)
	if err != nil {
		fmt.Println(err)
	}
	//building my slice
	slice := build_slice(z)

	//testing addition with for loop
	now := time.Now()
	sum := add_slice(slice)
	fmt.Println("Sum from for loop addition", sum)
	fmt.Println("Time to do with for loop addition", time.Since(now))

	//sourced from tour of go, concurrent addition
	c := make(chan int)
	now = time.Now()
	go summ(slice[:len(slice)/2], c)
	go summ(slice[len(slice)/2:], c)
	//receive from c
	x, y := <-c, <-c
	fmt.Println("Sum from concurrent addition", x, y, x+y)
	fmt.Println("Time to add with concurrency", time.Since(now))

}
