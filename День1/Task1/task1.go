package main

//Генератор квадратов натуральных чисел

import (
	"fmt"
)

const count int = 10

func main() {
	squares := GetSquares(count)
	fmt.Printf("Квадраты первых %d чисел: %v\n", count, squares)
}

func GetSquares(n int) [] int{
	m := make([]int, n)
	for i := 1; i <= n; i++ {
		m[i-1] = i * i
	}
	return m
}