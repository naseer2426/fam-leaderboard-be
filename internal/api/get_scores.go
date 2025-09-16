package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naseer2426/fam-leaderboard-be/internal/db"
)

// GetScores handles GET /scores to list all scores for a given game type.
// Query params:
// - gameType: required, the game type to filter scores by
func GetScores(c *gin.Context) {
	gameType := c.Query("gameType")
	if gameType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gameType query parameter is required"})
		return
	}

	database := db.GetDB()

	var scores []db.Scoreboard
	result := database.Where("game_type = ?", gameType).
		Find(&scores)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch scores"})
		return
	}

	c.JSON(http.StatusOK, scores)
}
