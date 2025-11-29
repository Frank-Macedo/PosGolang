package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// println(Divisors(4))
	// println(DNAStrand("ATTGC"))
	// println(DontGiveMeFive(4, 17))
	fmt.Println(SortNumbers([]int{1, 2, 3, 10, 5}))
}

func DNAStrand(dna string) string {

	r1 := strings.NewReplacer("A", "T", "T", "A", "C", "G", "G", "C")
	dna = r1.Replace(dna)

	return dna

}

func Divisors(n int) int {

	var qtd int = 0

	for i := 1; i <= n; i++ {
		res := n % i
		if res == 0 {
			qtd++
		}
	}
	return qtd

}

func FindNextSquare(sq int64) int64 {

	value := int64(math.Sqrt(float64(sq)))
	if value*value == sq {
		value++
		return value * value
	}
	return -1

}

func DontGiveMeFive(start int, end int) int {

	numbers := make([]int, 0)

	for i := start; i <= end; i++ {

		n := i
		for n > 0 {
			if n%10 == 0 {
				numbers = append(numbers, n)
			}
			n = n / 10
		}
	}

	return len(numbers)

}

type number struct {
	p int
}

func SortNumbers(numbers []int) []int {

	return BubbleSort(numbers)
}

func BubbleSort(numbers []int) []int {
	// cria uma cópia para não alterar o slice original
	sorted := make([]int, len(numbers))
	n := len(sorted)
	copy(sorted, numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {

				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}

	}

	return sorted
}
