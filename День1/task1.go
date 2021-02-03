package День1

//Генератор квадратов натуральных чисел

import (
	"fmt"
	"os"
)

func main1() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	squares := GetSquares(n)
	for _, i := range squares{
		fmt.Println(i)
	}
}

func GetSquares(n int) [] int{
	m := make([]int, n)
	for i := 1; i <= n; i++ {
		m[i-1] = i * i
	}
	return m
}