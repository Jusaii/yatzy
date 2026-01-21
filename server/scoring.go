package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var UserMap = map[string]*GameData{}

func InitUser(id string, name string) {
	var newUser GameData
	newUser.name = name
	for i := range newUser.diceArray {
		newUser.diceArray[i].value = 0
		newUser.diceArray[i].locked = false
	}

	UserMap[id] = &newUser
}

func itemSelector(c *gin.Context) {
	what := c.Param("what")
	whatInt := 0
	var req idStruct
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	val := 0

	switch what {
	case "one":
		val = calculateNumbers(req.Id, 1)
		whatInt = 0
	case "two":
		val = calculateNumbers(req.Id, 2)
		whatInt = 1
	case "three":
		val = calculateNumbers(req.Id, 3)
		whatInt = 2
	case "four":
		val = calculateNumbers(req.Id, 4)
		whatInt = 3
	case "five":
		val = calculateNumbers(req.Id, 5)
		whatInt = 4
	case "six":
		val = calculateNumbers(req.Id, 6)
		whatInt = 5
	case "pair":
		whatInt = 6
		fmt.Println("pair clicked")
	case "doublepair":
		whatInt = 7
		fmt.Println("doublepair clicked")
	case "triples":
		whatInt = 8
		fmt.Println("triples clicked")
	case "quadruples":
		whatInt = 9
		fmt.Println("quadruples clicked")
	case "smallstraight":
		whatInt = 10
		fmt.Println("smallstraight clicked")
	case "bigstraight":
		whatInt = 11
		fmt.Println("bigstraight clicked")
	case "fullhouse":
		whatInt = 12
		fmt.Println("fullhouse clicked")
	case "mixed":
		whatInt = 13
		fmt.Println("mixed clicked")
	case "quintuples":
		whatInt = 14
		fmt.Println("quintuples clicked")
	}

	UserMap[req.Id].scorearray[whatInt] = val
	resetDice(req.Id)
}

func calculateTotal(id string) int {
	scoreRow := UserMap[id].scorearray
	total := 0
	for i := range 6 {
		total += scoreRow[i]
	}

	// Check if bonus is given
	if total >= 63 {
		total += 50
	}

	for i := 6; i < 15; i++ {
		total += scoreRow[i]
	}

	return total
}

func getScoreRow(id string) string {
	scoreRow := UserMap[id].scorearray
	var stringarr [15]string

	for i := range scoreRow {
		stringarr[i] = strconv.Itoa(scoreRow[i])
	}
	row := strings.Join(stringarr[:], ";")

	return row
}
