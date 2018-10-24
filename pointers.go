package main

import "fmt"

func main() {
	numbers := &[7]int{2, 3, 5, 34, 2, 34, 76}
	fmt.Println(average(numbers))
}

func average(numbers *[7]int) float64 {
	var sum float64
	sum = 0
	for _, number := range numbers {
		sum = sum + float64(number)
	}
	return (sum / float64(len(numbers)))
}
