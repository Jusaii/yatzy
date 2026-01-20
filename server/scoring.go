package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type GameData struct {
	name       string
	scorearray [15]int
	diceArray  [5]DiceData
}

type DiceData struct {
	value  int
	locked bool
}

type NewScoreReq struct {
	Id     string `json:"id"`
	Target int    `json:"target"`
	Value  int    `json:"value"`
}

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

func UpdateScores(c *gin.Context) {
	var req NewScoreReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	UserMap[req.Id].scorearray[req.Target] = req.Value

	resetDice(req.Id)

	c.IndentedJSON(http.StatusOK, "Game started")
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
