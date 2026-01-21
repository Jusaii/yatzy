package main

func calculateNumbers(id string, val int) int {
	dicearr := &UserMap[id].diceArray
	sum := 0

	for i := range dicearr {
		if dicearr[i].value == val {
			sum += val
		}
	}

	return sum
}
