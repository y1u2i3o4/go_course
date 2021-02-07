package main

import (
	"fmt"
	"math/rand"
)

const length int = 10

func main() {
	array := make([]int, length)
	for i:=0; i < length; i++{
		array[i] = rand.Intn(10)
	}

	fmt.Printf("Массив: %d\n", array)
	fmt.Printf("Уникальных элементов: %d", GetUniqueValuesCount(array))
}

func GetUniqueValuesCount(array []int) int{
	dict := make(map[int]int)
	for _, i := range array{
		dict[i]++
	}

	count := 0
	for _, v := range dict{
		if v == 1 {
			count++
		}
	}
	return count
}