package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
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

func main() {

	//this line copied from go documenation to get the first arg
	firstArg := os.Args[1]
	z, err := strconv.Atoi(firstArg)
	if err != nil {
		fmt.Println(err)
	}
	//building my slice
	slice := build_slice(z)

	//making a copy of the slice that I built so sort stable does not recieve an already sorted slice
	copy_slice := make([]int, len(slice))
	copy(copy_slice, slice)

	//sort slice
	now := time.Now()
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	fmt.Println(time.Since(now))

	//sort stable
	now = time.Now()
	sort.SliceStable(copy_slice, func(i, j int) bool { return copy_slice[i] < copy_slice[j] })
	fmt.Println(time.Since(now))

}
