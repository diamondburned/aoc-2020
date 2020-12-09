package aocutil

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func SplitFile(name, split string) []string {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), split)
}

func MustAtoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func Window(numbers []int, size int, fn func([]int)) {
	if size > len(numbers) {
		return
	}

	for n := range numbers {
		if n+size > len(numbers) {
			break
		}
		fn(numbers[n : n+size])
	}
}

func Chunk(numbers []int, size int, fn func([]int)) {
	for i := 0; i < len(numbers); i += size {
		end := Min(i+size, len(numbers))
		fn(numbers[i:end])
	}
}

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func MinMaxes(numbers []int) (min, max int) {
	if len(numbers) == 0 {
		return -1, -1
	}

	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func Clamp(n, min, max int) int {
	if n > max {
		return max
	}
	if n < min {
		return min
	}
	return n
}
