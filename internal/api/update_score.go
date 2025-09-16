package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naseer2426/fam-leaderboard-be/internal/db"
	"gorm.io/gorm"
)

type updateScoreRequest struct {
	TeamName string `json:"teamName"`
	GameType string `json:"gameType"`
}

// IncreaseScore handles PUT /scores/increase_score to increment a team's score by 1.
func IncreaseScore(c *gin.Context) {
	var req updateScoreRequest
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
		UpdateColumn("team_score", gorm.Expr("team_score + ?", 1))

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

// DecreaseScore handles PUT /scores/decrease_score to decrement a team's score by 1.
func DecreaseScore(c *gin.Context) {
	var req updateScoreRequest
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
		UpdateColumn("team_score", gorm.Expr("team_score - ?", 1))

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
