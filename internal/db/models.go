package db

import "time"

// Scoreboard represents a team's score for a particular game type.
type Scoreboard struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	TeamName  string `gorm:"type:varchar(255);not null;index:idx_team_game,priority:1"`
	TeamScore int    `gorm:"not null"`
	GameType  string `gorm:"type:varchar(100);not null;index:idx_team_game,priority:2"`
}
