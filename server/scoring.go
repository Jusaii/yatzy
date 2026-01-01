package main

import (
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
)

type GameData struct {
	name       string
	scorearray [15]int
}

type NewScoreReq struct {
	Id     string `json:"id"`
	Target int    `json:"target"`
	Value  int    `json:"value"`
}

var UserMap = map[string]GameData{}

func InitUser(id string, name string) {
	var newUser GameData
	newUser.name = name

	UserMap[id] = newUser

	fmt.Println("New User created:")
	fmt.Println(UserMap[id].name)
	fmt.Println(UserMap[id].scorearray)
}

func UpdateScores(c *gin.Context) {
	var req NewScoreReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("New score update requested:")
	fmt.Println(req.Id, req.Target, req.Value)

	c.IndentedJSON(http.StatusOK, "Game started")
}
