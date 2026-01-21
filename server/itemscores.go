package main

import (
	"sort"
	"strconv"
)

func getDiceValuesArray(id string) [5]int {
	var valuesArray [5]int
	dicearr := &UserMap[id].diceArray
	for i := range valuesArray {
		valuesArray[i] = dicearr[i].value
	}
	sort.Ints(valuesArray[:])
	return valuesArray
}

func calculateNumbers(id string, val int) int {
	dicearr := getDiceValuesArray(id)
	sum := 0

	for i := range dicearr {
		if dicearr[i] == val {
			sum += val
		}
	}

	return sum
}

func calculatePair(id string) int {
	dicearr := getDiceValuesArray(id)
	sum := 0
	for i := 4; i >= 1; i-- {
		if dicearr[i] == dicearr[i-1] {
			info("pair found: " + strconv.Itoa(dicearr[i]))
			sum = 2 * dicearr[i]
			break
		}
	}

	return sum
}

func calculateDoublepair(id string) int {
	sum := 0
	info("doublepair clicked " + id)

	return sum
}

func calculateTriples(id string) int {
	sum := 0
	info("triples clicked " + id)

	return sum
}

func calculateQuadruples(id string) int {
	sum := 0
	info("quadruples clicked " + id)

	return sum
}

func calculateSmallstraight(id string) int {
	sum := 0
	info("small straight clicked " + id)

	return sum
}

func calculateBigstraight(id string) int {
	sum := 0
	info("big straight clicked " + id)

	return sum
}

func calculateFullhouse(id string) int {
	sum := 0
	info("full house clicked " + id)

	return sum
}

func calculateMixed(id string) int {
	sum := 0
	info("mixed clicked " + id)

	return sum
}

func calculateQuintuples(id string) int {
	sum := 0
	info("quintuples clicked " + id)

	return sum
}
