package main

func Sum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numsToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails 将切片尾部元素相加
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
