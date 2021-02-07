package main

import (
	"fmt"
	"math/rand"
)

const shift int = 3
const length int = 10

func main() {
	array := make([]int, length)
	for i:=0; i < length; i++{
		array[i] = rand.Intn(10)
	}
	fmt.Println("Original array", array)
	array = shiftArray(array, shift)
	fmt.Println("Shift: ", shift)
	fmt.Println("Shifted array", array)
}

func shiftArray(array []int, shift int) []int{
	len := len(array)
	tmp := make([]int, len)
	for i:=0; i < len; i++{
		tmp[(i+shift) % len] = array[i]
	}
	return tmp
}