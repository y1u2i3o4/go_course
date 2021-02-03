package День1

import (
	"fmt"
	"os"
)

//Найти самую длинную последовательность нулей
//func (N int) int

func main2() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	fmt.Printf("Число: %b, максимальная последовательность нулей %d",n, GetZeroSequenceLength(n))
}


func GetZeroSequenceLength(n int) int{
	var max, count = 0, 0
	for n > 0 {
		if n & 1 == 0 {
			count++
		} else {
			if max < count{
				max = count
			}
			count = 0
		}
		n >>= 1
	}
	if max > count{
		return  max
	} else {
		return count
	}
}