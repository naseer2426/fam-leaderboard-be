package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naseer2426/fam-leaderboard-be/internal/db"
)

type setScoreRequest struct {
	TeamName  string `json:"teamName"`
	GameType  string `json:"gameType"`
	TeamScore int    `json:"teamScore"`
}

// SetScore handles PUT /scores/set_score to set a team's score to a specific value.
func SetScore(c *gin.Context) {
	var req setScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON payload"})
		return
	}
	if req.TeamName == "" || req.GameType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teamName and gameType are required"})
		return
	}

	database := db.GetDB()

	result := database.Model(&db.Scoreboard{}).
		Where("team_name = ? AND game_type = ?", req.TeamName, req.GameType).
		Update("team_score", req.TeamScore)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update score"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "score not found for given teamName and gameType"})
		return
	}

	var updated db.Scoreboard
	if err := database.Where("team_name = ? AND game_type = ?", req.TeamName, req.GameType).First(&updated).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch updated score"})
		return
	}

	c.JSON(http.StatusOK, updated)
}
