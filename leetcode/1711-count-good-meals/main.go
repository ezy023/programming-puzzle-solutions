package main

import "math"

// Too slow
func countPairsOne(deliciousness []int) int {
	sums := make(map[int]int)

	for i := 0; i < len(deliciousness); i++ {
		for j := i + 1; j < len(deliciousness); j++ {
			s := deliciousness[i] + deliciousness[j]
			sums[s]++
		}
	}

	isPowerTwo := func(num int) bool {
		c := 0
		for num > 0 {
			if num&1 == 1 {
				c++
			}
			num = num >> 1
		}
		return c == 1
	}
	sum := 0
	for key, val := range sums {
		if isPowerTwo(key) {
			sum += val
		}
	}

	return sum
}

// The constraints of the problem state that each number in 'deliciousness'  will have an upper bound of 2^20
//   thus the most any sum could be is 2^20 + 2^20 = 2^21

// Too slow
func countPairsTwo(deliciousness []int) int {
	powersOfTwo := make(map[int]bool)

	for i := 0; i < 21; i++ {
		num := 1 << i
		powersOfTwo[num] = true
	}

	goodMeals := 0
	for i := 0; i < len(deliciousness); i++ {
		for j := i + 1; j < len(deliciousness); j++ {
			s := deliciousness[i] + deliciousness[j]
			if _, ok := powersOfTwo[s]; ok {
				goodMeals++
			}
		}
	}

	return goodMeals
}

func countPairs(deliciousness []int) int {
	powersOfTwo := make([]int, 22)

	for i := 0; i < 22; i++ {
		powersOfTwo[i] = 1 << i
	}

	numbers := make(map[int]int)
	modulo := int(math.Pow10(9) + 7)
	result := 0
	for i := 0; i < len(deliciousness); i++ {
		for j := 0; j < len(powersOfTwo); j++ {
			diff := powersOfTwo[j] - deliciousness[i]
			if diff >= 0 {
				result += numbers[diff]
				result = result % modulo
			}
		}
		numbers[deliciousness[i]]++
	}
	return result
}
