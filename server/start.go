package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type startGameRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func getReq(c *gin.Context) (startGameRequest, error) {
	var req startGameRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	return req, err
}

func startGame(c *gin.Context) {
	req, err := getReq(c)
	if err != nil {
		return
	}

	InitUser(req.Id, req.Name)

	c.IndentedJSON(http.StatusOK, "Game started")
}

func restartGame(c *gin.Context) {
	req, err := getReq(c)
	if err != nil {
		return
	}

	delete(UserMap, req.Id)

	c.IndentedJSON(http.StatusOK, "Game reset")
}
