package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type startGameRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func startGame(c *gin.Context) {
	var req startGameRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	InitUser(req.Id, req.Name)

	c.IndentedJSON(http.StatusOK, "Game started")
}
