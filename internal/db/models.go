package db

import "time"

// Scoreboard represents a team's score for a particular game type.
type Scoreboard struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	TeamName  string `gorm:"type:varchar(255);not null;index:idx_team_game,priority:1" json:"teamName"`
	TeamScore int    `gorm:"not null" json:"teamScore"`
	GameType  string `gorm:"type:varchar(100);not null;index:idx_team_game,priority:2" json:"gameType"`
}
