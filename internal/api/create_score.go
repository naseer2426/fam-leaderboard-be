package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naseer2426/fam-leaderboard-be/internal/db"
)

type createScoreRequest struct {
	TeamName  string `json:"teamName"`
	TeamScore int    `json:"teamScore"`
	GameType  string `json:"gameType"`
}

// CreateScore handles POST /scores to create a new Scoreboard entry.
func CreateScore(c *gin.Context) {
	var req createScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON payload"})
		return
	}

	if req.TeamName == "" || req.GameType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teamName and gameType are required"})
		return
	}

	database := db.GetDB()

	score := db.Scoreboard{
		TeamName:  req.TeamName,
		TeamScore: req.TeamScore,
		GameType:  req.GameType,
	}

	if err := database.Create(&score).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create score"})
		return
	}

	c.JSON(http.StatusCreated, score)
}
